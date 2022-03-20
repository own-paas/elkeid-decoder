package schema

var (
	Handles map[string]Schema
)

func init() {
	Handles = map[string]Schema{}
	Handles["42"] = &Connect{}
	Handles["49"] = &Bind{}
	Handles["59"] = &Execve{}
	Handles["82"] = &Rename{}
	Handles["165"] = &Mount{}
	Handles["356"] = &MemfdCreate{}
	Handles["601"] = &DnsQuery{}
	Handles["602"] = &CreateFile{}
	Handles["603"] = &LoadModule{}
	Handles["604"] = &UpdateCred{}
	Handles["607"] = &CallUserModeHelperExec{}
	Handles["610"] = &UsbDeviceEvent{}
	Handles["611"] = &PrivilegeEscalation{}
	Handles["700"] = &ProcFile{}
	Handles["701"] = &SyscallTableHook{}
	Handles["702"] = &HiddenKernelModule{}
	Handles["703"] = &InterruptTableHook{}
}
