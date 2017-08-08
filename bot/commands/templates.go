package commands

//failures
const pool_does_not_exist_claim = "What is %s?"
const pool_not_specified_claim = "I feel nothing"
const pool_already_claimed_claim = "%s is taken. In the end there can only be one."
const pool_does_not_exist_destroy = "What is %s?"
const pool_not_specified_destroy = "I feel nothing"
const pool_does_not_exist_owner = "What is %s?"
const pool_is_not_claimed_owner = "%s is not claimed. I have the power!"
const pool_is_not_claimed_release = "%s is not claimed. I have the power!"

const success_owner = "Patience, Highlander. %s was claimed by %s on %s"
const success_status = "I know... I know everything! I am everything!\n *Claimed by you:* %s\n*Claimed by others:* %s\n*Unclaimed:* %s"
const success_claim = "You have claimed %s. You have power beyond imagination. Use it well my friend. Don't lose your head."
const success_create = "%s has been created. From the dawn of time we came; moving silently down through the centuries, living many secret lives, struggling to reach the time of the Gathering; when the few who remain will battle to the last. No one has ever known we were among you... until now."
const success_help = "I am Connor MacLaim of the Clan MacLaim. I was born in 1518 in the village of Glenfinnan on the shores of Loch Shiel. And I am immortal." +
	"```\n" +
	"  claim <env> [<message>]   Claim an unclaimed environment\n" +
	"  create <env>              Create a new environment\n" +
	"  destroy <env>             Destroy an environment\n" +
	"  owner <env>               Show the user who claimed the environment\n" +
	"  release <env>             Release a claimed environment\n" +
	"  status                    Show claimed and unclaimed environments\n" +
	"  help                      Display this message\n" +
	"```"
const success_destroy = "Destroyed %s. It is better to burn out than fade away."
const success_release = "Released %s. I apologize for calling your wife a bloated warthog, and I bid you good day."
