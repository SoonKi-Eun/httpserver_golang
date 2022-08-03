# httpserver_golang
simple golang httpserver

$ cd go_source/bin
$ ./0_create_go_mod.sh
---------------------------------------------------------------------------------------------------
========================================================
Create module init(kbell) >>
go: creating new go.mod: module kbell
go: to add module requirements and sums:
        go mod tidy
========================================================
module show >>
file :  /home/esk1223/project_mng_dir/_Git_SkySilver/httpserver_golang/go_source/src/go.mod
module kbell

go 1.17
========================================================
========================================================
Project Path >>
: /home/esk1223/project_mng_dir/_Git_SkySilver/httpserver_golang/go_source
GOROOT >>
: /home/esk1223/env_default_pkg/go/pkg/go-1.17.5-linux-amd64
GOPATH >>
: /home/esk1223/project_mng_dir/_Git_SkySilver/httpserver_golang/go_source
========================================================
Proejct Go Run (ver. interpreter) >>
go: downloading github.com/sirupsen/logrus v1.9.0
go: downloading golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8
go get: added github.com/sirupsen/logrus v1.9.0
go get: added golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8
go: downloading gopkg.in/natefinch/lumberjack.v2 v2.0.0
go get: added gopkg.in/natefinch/lumberjack.v2 v2.0.0
go: downloading github.com/robfig/cron/v3 v3.0.1
go: downloading github.com/robfig/cron v1.2.0
go get: added github.com/robfig/cron/v3 v3.0.1
go: downloading github.com/gookit/config v1.1.0
go: downloading github.com/gookit/config/v2 v2.1.2
go: downloading github.com/gookit/goutil v0.5.2
go: downloading github.com/imdario/mergo v0.3.13
go: downloading github.com/mitchellh/mapstructure v1.5.0
go: downloading github.com/mattn/go-isatty v0.0.14
go: downloading github.com/mitchellh/go-homedir v1.1.0
go get: added github.com/gookit/config/v2 v2.1.2
go get: added github.com/gookit/goutil v0.5.2
go get: added github.com/imdario/mergo v0.3.13
go get: added github.com/mattn/go-isatty v0.0.14
go get: added github.com/mitchellh/go-homedir v1.1.0
go get: added github.com/mitchellh/mapstructure v1.5.0
go: downloading gopkg.in/yaml.v2 v2.4.0
go get: added github.com/robfig/cron v1.2.0
go: downloading github.com/labstack/echo v3.3.10+incompatible
go: downloading golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad
go: downloading golang.org/x/net v0.0.0-20200301022130-244492dfa37a
go: downloading golang.org/x/text v0.3.5
go: downloading github.com/labstack/gommon v0.3.1
go: downloading github.com/mattn/go-colorable v0.1.11
go: downloading github.com/valyala/fasttemplate v1.2.1
go: downloading github.com/valyala/bytebufferpool v1.0.0
go get: added github.com/labstack/echo v3.3.10+incompatible
go get: added github.com/labstack/gommon v0.3.1
go get: added github.com/mattn/go-colorable v0.1.11
go get: added github.com/valyala/bytebufferpool v1.0.0
go get: added github.com/valyala/fasttemplate v1.2.1
go: downloading github.com/labstack/echo/v4 v4.7.2
go: downloading golang.org/x/crypto v0.0.0-20210817164053-32db794688a5
go: downloading golang.org/x/net v0.0.0-20211015210444-4f30a5c0130f
go: downloading golang.org/x/text v0.3.7
go get: added github.com/labstack/echo/v4 v4.7.2
go get: upgraded golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad => v0.0.0-20210817164053-32db794688a5
go get: upgraded golang.org/x/net v0.0.0-20200301022130-244492dfa37a => v0.0.0-20211015210444-4f30a5c0130f
go get: upgraded golang.org/x/text v0.3.5 => v0.3.7
go: downloading github.com/golang-jwt/jwt v3.2.2+incompatible
go: downloading golang.org/x/time v0.0.0-20201208040808-7e3f01d25324
---------------------------------------------------------------------------------------------------

$./1_src_run.sh
---------------------------------------------------------------------------------------------------
========================================================
Project Path >>
: /home/esk1223/project_mng_dir/_Git_SkySilver/httpserver_golang/go_source
GOROOT >>
: /home/esk1223/env_default_pkg/go/pkg/go-1.17.5-linux-amd64
GOPATH >>
: /home/esk1223/project_mng_dir/_Git_SkySilver/httpserver_golang/go_source
========================================================
Proejct Go Run (ver. interpreter) >>
[    INFO :: 2022-08-03 10:40:49 ] ::: ==============================================================
[    INFO :: 2022-08-03 10:40:49 ] ::: project >>>>>>>>>>
[    INFO :: 2022-08-03 10:40:49 ] :::  total_cpu               : 8
[    INFO :: 2022-08-03 10:40:49 ] :::  porject_name            : system_controller
[    INFO :: 2022-08-03 10:40:49 ] :::  porject_path            : /home/esk1223/SDWAN/system_controller_golang/go_source
[    INFO :: 2022-08-03 10:40:49 ] ::: http    >>>>>>>>>>
[    INFO :: 2022-08-03 10:40:49 ] :::  port                    : 18001
[    INFO :: 2022-08-03 10:40:49 ] :::  timeout                 : 10ns
[    INFO :: 2022-08-03 10:40:49 ] :::  ssl_active              : false
[    INFO :: 2022-08-03 10:40:49 ] :::  private_key_path        : /home/esk1223/SDWAN/system_controller_golang/go_source/key/server.key
[    INFO :: 2022-08-03 10:40:49 ] :::  crt_certificate_path    : /home/esk1223/SDWAN/system_controller_golang/go_source/key/server.crt
[    INFO :: 2022-08-03 10:40:49 ] ::: ==============================================================
[    INFO :: 2022-08-03 10:40:49 ] ::: lib_module package init >>>
[    INFO :: 2022-08-03 10:40:49 ] ::: task_module package init >>>
[    INFO :: 2022-08-03 10:40:49 ] ::: http_module package init >>>
[    INFO :: 2022-08-03 10:40:49 ] ::: main package init >>>
---------------------------------------------------------------------------------------------------

$curl 192.168.1.108:18001
---------------------------------------------------------------------------------------------------
sys_controller active
---------------------------------------------------------------------------------------------------
