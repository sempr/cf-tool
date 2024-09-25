package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	ansi "github.com/k0kubun/go-ansi"
	"github.com/mitchellh/go-homedir"
	"github.com/sempr/cf/client"
	"github.com/sempr/cf/cmd"
	"github.com/sempr/cf/config"

	docopt "github.com/docopt/docopt-go"
)

var version = "v1.0.0"
const configPath = "~/.atcoder/config"
const sessionPath = "~/.atcoder/session"

func main() {
	usage := `AtCoder Tool $%version%$ (atc). https://github.com/sempr/cf

You should run "$%CF%$ config" to configure your handle, password and code
templates at first.

If you want to compete, the best command is "$%CF%$ race"

Usage:
  $%CF%$ config
  $%CF%$ submit [-f <file>] [<specifier>...]
  $%CF%$ list [<specifier>...]
  $%CF%$ parse [<specifier>...]
  $%CF%$ gen [<alias>]
  $%CF%$ test [<file>]
  $%CF%$ watch [all] [<specifier>...]
  $%CF%$ open [<specifier>...]
  $%CF%$ stand [<specifier>...]
  $%CF%$ sid [<specifier>...]
  $%CF%$ race [<specifier>...]
  $%CF%$ pull [ac] [<specifier>...]
  $%CF%$ clone [ac] [<handle>]
  $%CF%$ upgrade

Options:
  -h --help            Show this screen.
  --version            Show version.
  -f <file>, --file <file>, <file>
                       Path to file. E.g. "a.cpp", "./temp/a.cpp"
  <specifier>          Any useful text. E.g.
                       "https://atcoder.jp/contests/abc259",
                       "https://atcoder.jp/contests/abc259/tasks/abc259_a",
                       "1111A", "1111", "a", "Cw4JRyRGXR"
                       You can combine multiple specifiers to specify what you
                       want.
  <alias>              Template's alias. E.g. "cpp"
  ac                   The status of the submission is Accepted.

Examples:
  $%CF%$ config            Configure the cf-tool.
  $%CF%$ submit            $%CF%$ will detect what you want to submit automatically.
  $%CF%$ submit -f a.cpp
  $%CF%$ submit https://atcoder.jp/contests/abc259/tasks/abc259_a
  $%CF%$ submit -f a.cpp 100A 
  $%CF%$ submit -f a.cpp 100 a
  $%CF%$ submit contest 100 a
  $%CF%$ submit gym 100001 a
  $%CF%$ list              List all problems' stats of a contest.
  $%CF%$ list 1119
  $%CF%$ parse 100         Fetch all problems' samples of contest 100 into
                       "{cf}/{contest}/100/<problem-id>".
  $%CF%$ parse gym 100001a
                       Fetch samples of problem "a" of gym 100001 into
                       "{cf}/{gym}/100001/a".
  $%CF%$ parse gym 100001
                       Fetch all problems' samples of gym 100001 into
                       "{cf}/{gym}/100001".
  $%CF%$ parse             Fetch samples of current problem into current path.
  $%CF%$ gen               Generate a code from default template.
  $%CF%$ gen cpp           Generate a code from the template whose alias is "cpp"
                       into current path.
  $%CF%$ test              Run the commands of a template in current path. Then
                       test all samples. If you want to add a new testcase,
                       create two files "inK.txt" and "ansK.txt" where K is
                       a string with 0~9.
  $%CF%$ watch             Watch the first 10 submissions of current contest.
  $%CF%$ watch all         Watch all submissions of current contest.
  $%CF%$ open 1136a        Use default web browser to open the page of contest
                       1136, problem a.
  $%CF%$ open gym 100136   Use default web browser to open the page of gym
                       100136.
  $%CF%$ stand             Use default web browser to open the standing page.
  $%CF%$ sid 52531875      Use default web browser to open the submission
                       52531875's page.
  $%CF%$ sid               Open the last submission's page.
  $%CF%$ race 1136         If the contest 1136 has not started yet, it will
                       countdown. When the countdown ends, it will open all
                       problems' pages and parse samples.
  $%CF%$ pull 100          Pull all problems' latest codes of contest 100 into
                       "./100/<problem-id>".
  $%CF%$ pull 100 a        Pull the latest code of problem "a" of contest 100 into
                       "./100/<problem-id>".
  $%CF%$ pull ac 100 a     Pull the "Accepted" or "Pretests passed" code of problem
                       "a" of contest 100.
  $%CF%$ pull              Pull the latest codes of current problem into current
                       path.
  $%CF%$ clone xalanq      Clone all codes of xalanq.
  $%CF%$ upgrade           Upgrade the "cf" to the latest version from GitHub.

File:
  $%CF%$ will save some data in some files:

  "~/.atcoder/config"        Configuration file, including templates, etc.
  "~/.atcoder/session"       Session file, including cookies, handle, password, etc.

  "~" is the home directory of current user in your system.

Template:
  You can insert some placeholders into your template code. When generate a code
  from the template, $%CF%$ will replace all placeholders by following rules:

  $%U%$   Handle (e.g. xalanq)
  $%Y%$   Year   (e.g. 2019)
  $%M%$   Month  (e.g. 04)
  $%D%$   Day    (e.g. 09)
  $%h%$   Hour   (e.g. 08)
  $%m%$   Minute (e.g. 05)
  $%s%$   Second (e.g. 00)

Script in template:
  Template will run 3 scripts in sequence when you run "$%CF%$ test":
    - before_script   (execute once)
    - script          (execute the number of samples times)
    - after_script    (execute once)
  You could set "before_script" or "after_script" to empty string, meaning
  not executing.
  You have to run your program in "script" with standard input/output (no
  need to redirect).

  You can insert some placeholders in your scripts. When execute a script,
  $%CF%$ will replace all placeholders by following rules:

  $%path%$   Path to source file (Excluding $%full%$, e.g. "/home/xalanq/")
  $%full%$   Full name of source file (e.g. "a.cpp")
  $%file%$   Name of source file (Excluding suffix, e.g. "a")
  $%rand%$   Random string with 8 character (including "a-z" "0-9")`
	color.Output = ansi.NewAnsiStdout()

	usage = strings.Replace(usage, `$%version%$`, version, 1)
	usage = strings.ReplaceAll(usage, `$%CF%$`, os.Args[0])
	opts, _ := docopt.ParseArgs(usage, os.Args[1:], fmt.Sprintf("AtCoder Tool (cf) %v", version))
	opts[`{version}`] = version

	cfgPath, _ := homedir.Expand(configPath)
	clnPath, _ := homedir.Expand(sessionPath)
	config.Init(cfgPath)
	client.Init(clnPath, config.Instance.Host, config.Instance.Proxy)

	err := cmd.Eval(opts)
	if err != nil {
		color.Red(err.Error())
	}
	color.Unset()
}
