/*
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     https://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package electron

import (
	"fmt"
	"os/exec"
	"syscall"
)

// RunElectron executes the standalone electron app which serves up our web components
func RunElectron() error {
	cmd := exec.Command("npm", "start", "--prefix", "web")

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("electron %d", err.(*exec.ExitError).Sys().(syscall.WaitStatus).ExitStatus())
	}

	return nil
}
