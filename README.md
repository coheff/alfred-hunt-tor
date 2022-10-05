# alfred-hunt-tor
An Alfred torrent search workflow powered by Go

# Download & installation
Grab the the latest version from the [releases page](https://github.com/coheff/alfred-hunt-tor/releases/tag/v0.1.0). Double click workflow file to import into Alfred.

If running on macOS Catalina or later, you _**MUST**_ add Alfred to the list of security exceptions for running unsigned software. Step-by-step instructions are available on the awgo wiki [here](https://github.com/deanishe/awgo/wiki/Catalina).

You can also start your application once ("with terminal") and except to open it. This way, Alfred will not be able to run just anything. To achieve it, you can install the workflow and then right click it to open the folder in Finder. Then right-click the 'alfred-hunt-tor' and open with Terminal. Then agree to open it. It will run and complain, not being started by Alfred. After that, the workflow will work (until an update of the executable). (Thanks [@oderwat](https://github.com/coheff/alfred-hunt-tor/issues/1#issuecomment-1268206311))

# Usage
Trigger a search using the keyword `t` followed by a search query. \
Seach categories using `#movies`, `#music`, and `#tv` tags. \
Order results using `-s` (seeders) and `-t` (time) flags. Ordering is descending only, for now.

# License
Distributed under the MIT License. See [LICENSE](https://github.com/coheff/alfred-hunt-tor/blob/main/LICENSE) for more information.
