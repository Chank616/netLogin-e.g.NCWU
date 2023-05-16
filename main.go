package main

import (
	"fmt"
	"netLogin/utils"
	"regexp"
	"time"

	"github.com/go-toast/toast"
)

func main() {
	username := "16696682230"
	password := "654321"

	// ip
	params := make(map[string]string)
	now := fmt.Sprintf("%d", time.Now().Unix()*100)
	params["callback"] = "jQuery112406336369815771166_" + now
	params["_"] = now
	resp := utils.Get("http://192.168.0.170/cgi-bin/rad_user_info", params)
	compileRegex := regexp.MustCompile(`online_ip":"(.*?)",`)
	matchArr := compileRegex.FindStringSubmatch(resp)
	ip := matchArr[len(matchArr)-1]

	// challenge
	params["ip"] = ip
	params["username"] = username
	resp = utils.Get("http://192.168.0.170/cgi-bin/get_challenge", params)
	compileRegex = regexp.MustCompile(`challenge":"(.*?)",`)
	matchArr = compileRegex.FindStringSubmatch(resp)
	challenge := matchArr[len(matchArr)-1]

	// password参数
	passwordEncrypt := utils.Hmac(challenge, password)

	// info参数
	i := fmt.Sprintf(`{"username":"%s","password":"%s","ip":"%s","acid":"2","enc_ver":"srun_bx1"}`, username, password, ip)
	infoEncrypt := "{SRBX1}" + utils.Base64Encode(utils.GetXencode(i, challenge))

	// chksum参数
	token := challenge
	chkstr := token + username
	chkstr += token + passwordEncrypt
	chkstr += token + "2"
	chkstr += token + ip
	chkstr += token + "200"
	chkstr += token + "1"
	chkstr += token + infoEncrypt
	chksumEncrypt := utils.Sha1(chkstr)
	
	// 发送登录包
	params["action"] = "login"
	params["password"] = "{MD5}" + passwordEncrypt
	params["ac_id"] = "2"
	params["chksum"] = chksumEncrypt
	params["info"] = infoEncrypt
	params["n"] = "200"
	params["type"] = "1"
	params["os"] = "Windows 10"
	params["name"] = "Windows"
	params["double_stack"] = "0"
	resp = utils.Get("http://192.168.0.170/cgi-bin/srun_portal", params)
	
	// 显示通知
	compileRegex = regexp.MustCompile(`suc_msg":"(.*?)",`)
	matchArr = compileRegex.FindStringSubmatch(resp)
	suc_msg := matchArr[len(matchArr)-1]
	notification := toast.Notification{
		AppID:   "Microsoft.Windows.Shell.RunDialog",
		Title:   "华北水利水电大学校园网认证状态",
		Message: suc_msg,
	}
	notification.Push()
}
