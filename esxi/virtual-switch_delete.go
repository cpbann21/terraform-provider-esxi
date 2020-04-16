package esxi

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceVirtualSwitchDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*Config)
	esxiSSHinfo := SshConnectionStruct{c.esxiHostName, c.esxiHostPort, c.esxiUserName, c.esxiPassword}
	log.Println("[resourceVirtualSwitchDelete]")

	var remote_cmd, cmd_result string
	var err error

	virtual_switch_name := d.Get("virtual_switch_name").(string)

	remote_cmd = fmt.Sprintf(`vim-cmd hostsvc/net/vswitch_remove "%s"`, virtual_switch_name)
	cmd_result, err = runRemoteSshCommand(esxiSSHinfo, remote_cmd, "destroy virtual switch")

	if err != nil {
		return fmt.Errorf("Unable to destroy virtual switch: %s", err)
	}

	if cmd_result != "" {
		return fmt.Errorf("Unable to destroy virtual switch: %s", cmd_result)
	}

	d.SetId("")
	return nil
}
