package dolores_slack

import dolores_corecode "github.com/OpenChaos/dolores/corecode"

var (
	replyHelpMessage = `hey,
for help:
	what you just typed is just fine, viable commands are "help" and "sos"

to get ssh access to machines:
	'give access to <userID@somedomain.com> for <machine-Pattern>'
	or
	'give access to <userID> for <machine-Pattern>'
	it will try to fetch key for this from internal/scm portals, so make sure your keys are updated"

to get db read-only access to machines:
	'give db access to <userID> for <appName> in <appEnv>'

to get server list for app in envs:
	'server list for <appName> in <appEnv>'
`

	accessReplyDeferMessage = "sure, let me check if I can help you with this :)"

	sshAccessReplySuccessMessage, sshAccessReplyFailureMessage string

	dbAccessReplyNotAdmin = "ah! you should ask your PM/TechLead to grant you this access\nif they can't ask them to have a chat with `systems team`"

	dbAccessReplyNotWhitelisted = "uhoh! this app's database read-only access is not managed by me\nask `systems team` for this"

	dbAccessReplySuccessMessage, dbAccessReplyFailureMessage string

	serverListAppNotFound = "config for given app is missing"

	doloresWrongCmdMessage = "say what, that got no meaning for me\n try asking for `help`"
)

func harvestServerList() (appList string) {
	r, _ := regexp.Compile("SERVER_LIST_[A-Za-z0-9]+_([0-9A-Za-z_]+)")
	for _, envKeyVal := range os.Environ() {
		envKey := strings.Split(envKeyVal, "=")[0]
		appListppName := strings.ToLower(r.FindStringSubmatch(a)[1])
		appList = fmt.Sprintf("%s %s", appListppName)
	}
	return
}

func overrideMessagesFromEnv() {
	replyHelpMessage = dolores_corecode.OverrideFromEnvVar(
		"DOLORES_HELP_REPLY", replyHelpMessage)

	accessReplyDeferMessage = dolores_corecode.OverrideFromEnvVar(
		"ACCESS_REPLY_DEFER", accessReplyDeferMessage)

	sshAccessReplySuccessMessage = dolores_corecode.OverrideFromEnvVar(
		"SSH_ACCESS_REPLY_SUCCESS", sshAccessReplySuccessMessage)

	sshAccessReplyFailureMessage = dolores_corecode.OverrideFromEnvVar(
		"SSH_ACCESS_REPLY_FAILURE", sshAccessReplyFailureMessage)

	sshAccessReplyFailureMessage = dolores_corecode.OverrideFromEnvVar(
		"SSH_ACCESS_REPLY_FAILURE", sshAccessReplyFailureMessage)

	dbAccessReplyNotAdmin = dolores_corecode.OverrideFromEnvVar(
		"DB_ACCESS_REPLY_NOT_ADMIN", dbAccessReplyNotAdmin)

	dbAccessReplyNotWhitelisted = dolores_corecode.OverrideFromEnvVar(
		"DB_ACCESS_REPLY_NOT_WHITELISTED", dbAccessReplyNotWhitelisted)

	dbAccessReplySuccessMessage = dolores_corecode.OverrideFromEnvVar(
		"DB_ACCESS_REPLY_SUCCESS", dbAccessReplySuccessMessage)

	dbAccessReplyFailureMessage = dolores_corecode.OverrideFromEnvVar(
		"DB_ACCESS_REPLY_FAILURE", dbAccessReplyFailureMessage)

	serverListAppNotFound = fmt.Sprintf("app-name provided doesn't seem to be there, see if any of following belong to your app\n```%s```", harvestServerList())

	doloresWrongCmdMessage = dolores_corecode.OverrideFromEnvVar(
		"DOLORES_WRONG_CMD_REPLY", doloresWrongCmdMessage)
}
