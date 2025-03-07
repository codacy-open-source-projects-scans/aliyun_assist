// Copyright (c) 2009-present, Alibaba Cloud All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package config

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/aliyun/aliyun_assist_client/agent/session/plugin/i18n"

	"github.com/aliyun/aliyun_assist_client/agent/session/plugin/cli"
)

var hookLoadConfiguration = func(fn func(path string) (*Configuration, error)) func(path string) (*Configuration, error) {
	return fn
}

var hookSaveConfiguration = func(fn func(config *Configuration) error) func(config *Configuration) error {
	return fn
}

func NewConfigureCommand() *cli.Command {

	c := &cli.Command{
		Name: "configure",
		Short: i18n.T(
			"configure credential and settings",
			"配置身份认证和其他信息"),
		Usage: "configure --mode {AK|StsToken|CredentialsURI} --profile <profileName>",
		//	Usage: "configure --mode {AK|StsToken|RamRoleArn|EcsRamRole|RsaKeyPair|RamRoleArnWithRoleName} --profile <profileName>",
		Run: func(ctx *cli.Context, args []string) error {
			if len(args) > 0 {
				return cli.NewInvalidCommandError(args[0], ctx)
			}
			profileName, _ := ProfileFlag(ctx.Flags()).GetValue()
			mode, _ := ModeFlag(ctx.Flags()).GetValue()
			return doConfigure(ctx, profileName, mode)
		},
	}

	c.AddSubCommand(NewConfigureGetCommand())
	c.AddSubCommand(NewConfigureSetCommand())
	c.AddSubCommand(NewConfigureListCommand())
	c.AddSubCommand(NewConfigureDeleteCommand())
	return c
}

func doConfigure(ctx *cli.Context, profileName string, mode string) error {
	w := ctx.Writer()

	conf, err := hookLoadConfiguration(LoadConfiguration)(GetConfigPath2())
	if err != nil {
		return err
	}

	if profileName == "" {
		if conf.CurrentProfile == "" {
			profileName = "default"
		} else {
			profileName = conf.CurrentProfile
			originMode := string(conf.GetCurrentProfile(ctx).Mode)
			if mode == "" {
				mode = originMode
			} else if mode != originMode {
				cli.Printf(w, "Warning: You are changing the authentication type of profile '%s' from '%s' to '%s'\n", profileName, originMode, mode)
			}
		}
	}
	if mode == "" {
		mode = "AK"
	}
	cp, ok := conf.GetProfile(profileName)
	if !ok {
		cp = conf.NewProfile(profileName)
	}

	cli.Printf(w, "Configuring profile '%s' in '%s' authenticate mode...\n", profileName, mode)

	if mode != "" {
		switch AuthenticateMode(mode) {
		case AK:
			cp.Mode = AK
			configureAK(w, &cp, ctx)
		case StsToken:
			cp.Mode = StsToken
			configureStsToken(w, &cp, ctx)
		case RamRoleArn:
			cp.Mode = RamRoleArn
			configureRamRoleArn(w, &cp, ctx)
		case EcsRamRole:
			cp.Mode = EcsRamRole
			configureEcsRamRole(w, &cp)
		case RamRoleArnWithEcs:
			cp.Mode = RamRoleArnWithEcs
			configureRamRoleArnWithEcs(w, &cp)
		case RsaKeyPair:
			cp.Mode = RsaKeyPair
			configureRsaKeyPair(w, &cp)
		case External:
			cp.Mode = External
			configureExternal(w, &cp)
		case CredentialsURI:
			cp.Mode = CredentialsURI
			configureCredentialsURI(w, &cp)
		default:
			return fmt.Errorf("unexcepted authenticate mode: %s", mode)
		}
	} else {
		configureAK(w, &cp, ctx)
	}

	ak, _ := AccessKeyIdFlag(ctx.Flags()).GetValue()
	if ak == "" {
		// configure common
		cli.Printf(w, "Default Region Id [%s]: ", cp.RegionId)
		cp.RegionId = ReadInput(cp.RegionId)
		cli.Printf(w, "Default Output Format [%s]: json (Only support json)\n", cp.OutputFormat)

		// cp.OutputFormat = ReadInput(cp.OutputFormat)
		cp.OutputFormat = "json"

		cli.Printf(w, "Default Language [zh|en] %s: ", cp.Language)

		cp.Language = ReadInput(cp.Language)
		if cp.Language != "zh" && cp.Language != "en" {
			cp.Language = "en"
		}
	} else {
		region, _ := RegionFlag(ctx.Flags()).GetValue()
		if region != "" {
			cp.RegionId = region
		}
	}

	//fmt.Printf("User site: [china|international|japan] %s", cp.Site)
	//cp.Site = ReadInput(cp.Site)

	cli.Printf(w, "Saving profile[%s] ...", profileName)

	conf.PutProfile(cp)
	conf.CurrentProfile = cp.Name
	err = hookSaveConfiguration(SaveConfiguration)(conf)

	if err != nil {
		return err
	}
	cli.Printf(w, "Done.\n")

	// DoHello(ctx, &cp)
	return nil
}

func configureCredentialsURI(w io.Writer, cp *Profile) error {
	cli.Printf(w, "Credentials URI [%s]: ", cp.CredentialsURI)
	cp.CredentialsURI = ReadInput(cp.CredentialsURI)
	return nil
}

func configureAK(w io.Writer, cp *Profile, ctx *cli.Context) error {
	ak, _ := AccessKeyIdFlag(ctx.Flags()).GetValue()
	if ak != "" {
		cp.AccessKeyId = ak
		sk, _ := AccessKeySecretFlag(ctx.Flags()).GetValue()
		cp.AccessKeySecret = sk
		return nil
	}
	cli.Printf(w, "Access Key Id [%s]: ", MosaicString(cp.AccessKeyId, 3))
	cp.AccessKeyId = ReadInput(cp.AccessKeyId)
	cli.Printf(w, "Access Key Secret [%s]: ", MosaicString(cp.AccessKeySecret, 3))
	cp.AccessKeySecret = ReadInput(cp.AccessKeySecret)
	return nil
}

func configureStsToken(w io.Writer, cp *Profile, ctx *cli.Context) error {
	err := configureAK(w, cp, ctx)
	if err != nil {
		return err
	}
	token, _ := StsTokenFlag(ctx.Flags()).GetValue()
	if token != "" {
		cp.StsToken = token
		return nil
	}
	cli.Printf(w, "Sts Token [%s]: ", cp.StsToken)
	cp.StsToken = ReadInput(cp.StsToken)
	return nil
}

func configureRamRoleArn(w io.Writer, cp *Profile, ctx *cli.Context) error {
	err := configureAK(w, cp, ctx)
	if err != nil {
		return err
	}
	cli.Printf(w, "Sts Region [%s]: ", cp.StsRegion)
	cp.StsRegion = ReadInput(cp.StsRegion)
	cli.Printf(w, "Ram Role Arn [%s]: ", cp.RamRoleArn)
	cp.RamRoleArn = ReadInput(cp.RamRoleArn)
	cli.Printf(w, "Role Session Name [%s]: ", cp.RoleSessionName)
	cp.RoleSessionName = ReadInput(cp.RoleSessionName)
	if cp.ExpiredSeconds == 0 {
		cp.ExpiredSeconds = 900
	}
	cli.Printf(w, "Expired Seconds [%v]: ", cp.ExpiredSeconds)
	cp.ExpiredSeconds, _ = strconv.Atoi(ReadInput(strconv.Itoa(cp.ExpiredSeconds)))
	return nil
}

func configureEcsRamRole(w io.Writer, cp *Profile) error {
	cli.Printf(w, "Ecs Ram Role [%s]: ", cp.RamRoleName)
	cp.RamRoleName = ReadInput(cp.RamRoleName)
	return nil
}

func configureRamRoleArnWithEcs(w io.Writer, cp *Profile) error {
	cli.Printf(w, "Ecs Ram Role [%s]: ", cp.RamRoleName)
	cp.RamRoleName = ReadInput(cp.RamRoleName)
	cli.Printf(w, "Ram Role Arn [%s]: ", cp.RamRoleArn)
	cp.RamRoleArn = ReadInput(cp.RamRoleArn)
	cli.Printf(w, "Role Session Name [%s]: ", cp.RoleSessionName)
	cp.RoleSessionName = ReadInput(cp.RoleSessionName)
	if cp.ExpiredSeconds == 0 {
		cp.ExpiredSeconds = 900
	}
	cli.Printf(w, "Expired Seconds [%v]: ", cp.ExpiredSeconds)
	cp.ExpiredSeconds, _ = strconv.Atoi(ReadInput(strconv.Itoa(cp.ExpiredSeconds)))
	return nil
}

func configureRsaKeyPair(w io.Writer, cp *Profile) error {
	cli.Printf(w, "Rsa Private Key File: ")
	keyFile := ReadInput("")
	buf, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return fmt.Errorf("read key file %s failed %v", keyFile, err)
	}
	cp.PrivateKey = string(buf)
	cli.Printf(w, "Rsa Key Pair Name: ")
	cp.KeyPairName = ReadInput("")
	cp.ExpiredSeconds = 900
	return nil
}

func configureExternal(w io.Writer, cp *Profile) error {
	cli.Printf(w, "Process Command [%s]: ", cp.ProcessCommand)
	cp.ProcessCommand = ReadInput(cp.ProcessCommand)
	return nil
}

func ReadInput(defaultValue string) string {
	var s string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		s = scanner.Text()
	}
	if s == "" {
		return defaultValue
	}
	return s
}

func MosaicString(s string, lastChars int) string {
	r := len(s) - lastChars
	if r > 0 {
		return strings.Repeat("*", r) + s[r:]
	} else {
		return strings.Repeat("*", len(s))
	}
}

func GetLastChars(s string, lastChars int) string {
	r := len(s) - lastChars
	if r > 0 {
		return s[r:]
	} else {
		return strings.Repeat("*", len(s))
	}
}
