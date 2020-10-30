package libs

import (
	cloudwatch "github.com/yongshengli/gopher-lua-libs/aws/cloudwatch"
	cert_util "github.com/yongshengli/gopher-lua-libs/cert_util"
	chef "github.com/yongshengli/gopher-lua-libs/chef"
	cmd "github.com/yongshengli/gopher-lua-libs/cmd"
	crypto "github.com/yongshengli/gopher-lua-libs/crypto"
	db "github.com/yongshengli/gopher-lua-libs/db"
	filepath "github.com/yongshengli/gopher-lua-libs/filepath"
	goos "github.com/yongshengli/gopher-lua-libs/goos"
	http "github.com/yongshengli/gopher-lua-libs/http"
	humanize "github.com/yongshengli/gopher-lua-libs/humanize"
	inspect "github.com/yongshengli/gopher-lua-libs/inspect"
	ioutil "github.com/yongshengli/gopher-lua-libs/ioutil"
	json "github.com/yongshengli/gopher-lua-libs/json"
	log "github.com/yongshengli/gopher-lua-libs/log"
	plugin "github.com/yongshengli/gopher-lua-libs/plugin"
	pprof "github.com/yongshengli/gopher-lua-libs/pprof"
	prometheus "github.com/yongshengli/gopher-lua-libs/prometheus/client"
	regexp "github.com/yongshengli/gopher-lua-libs/regexp"
	runtime "github.com/yongshengli/gopher-lua-libs/runtime"
	"github.com/yongshengli/gopher-lua-libs/stats"
	storage "github.com/yongshengli/gopher-lua-libs/storage"
	strings "github.com/yongshengli/gopher-lua-libs/strings"
	tac "github.com/yongshengli/gopher-lua-libs/tac"
	tcp "github.com/yongshengli/gopher-lua-libs/tcp"
	telegram "github.com/yongshengli/gopher-lua-libs/telegram"
	template "github.com/yongshengli/gopher-lua-libs/template"
	time "github.com/yongshengli/gopher-lua-libs/time"
	xmlpath "github.com/yongshengli/gopher-lua-libs/xmlpath"
	yaml "github.com/yongshengli/gopher-lua-libs/yaml"
	zabbix "github.com/yongshengli/gopher-lua-libs/zabbix"

	lua "github.com/yuin/gopher-lua"
)

// Preload preload all gopher lua packages
func Preload(L *lua.LState) {
	time.Preload(L)
	strings.Preload(L)
	filepath.Preload(L)
	ioutil.Preload(L)
	http.Preload(L)
	regexp.Preload(L)
	tac.Preload(L)
	inspect.Preload(L)
	yaml.Preload(L)
	plugin.Preload(L)
	cmd.Preload(L)
	json.Preload(L)
	tcp.Preload(L)
	xmlpath.Preload(L)
	db.Preload(L)
	cert_util.Preload(L)
	runtime.Preload(L)
	telegram.Preload(L)
	zabbix.Preload(L)
	pprof.Preload(L)
	prometheus.Preload(L)
	crypto.Preload(L)
	goos.Preload(L)
	storage.Preload(L)
	humanize.Preload(L)
	chef.Preload(L)
	template.Preload(L)
	cloudwatch.Preload(L)
	log.Preload(L)
	stats.Preload(L)
}
