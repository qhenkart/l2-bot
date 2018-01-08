# L2 Revolution Bot CLI Tool

### Purpose
Lineage2 Revolution is a great game, but the game mechanics can be a bit repetitive and boring. Everyday we are forced to stare at our screens watching our character auto-battle for 1-2 hours before we actually get to play. Skipping the boring stuff seriously nerfs the character. This app was created to automate the boring aspects so when you wake up, you can get right to the fun stuff

### Risk
Using Macros, third party programs, emulators all violates the ToS, and while NetMarble is clearly not enforcing these things, using an app like this would be a bannable offense. *That being said*, this is a very low risk app to run. This tool does not make API calls, there is no hacking or package sniffing. It merely replicates user interaction with the UI using the SikuliX UI tool. From an administrative perspective, there is no way to detect the difference between a user pressing a button or this program doing it for you. However, it would be prudent not to talk about using this tool.

### Feature list
1. Runs all weeklys, sub, dailys, Temple Dungeon(Hard) and Summoning Dungeon(Normal) at Midnight, and finishes with grinding nearby mobs
2. Sends notifications to your phone —— must set it up via [PushOver](http://pushover.net)
3. Can run individual script using the CLI, including Main Quests (for leveling alts or level cap increases)
4. All scripts handle starting Nox, restarting Nox in the event of an error, logging in, disconnections, and anomolies such as pop-ups or event advertisements
5. Handles errors gracefully in that each script will be attempted 3 times, before moving to the next one, and if all scripts have errors, Nox will restart and attempt it again.
6. Outputs very verbose logs that you can connect to dropbox, icloud or google drive to view the live logs on your phone at any time. Outputs as log.txt

# Usage

### Setup
A Mac is required to run these scripts, although set-up is certainly possible on windows or a VM, I will leave it up to contributors to add instructions for Windows installation

1. Download [SikuliX](http://www.sikuli.org/)
2. Install Go (this is the most complicated part, there are lots of resources online to help if you get stuck)
```
  1) Create Directories (type this into your terminal)
    mkdir $HOME/Go
    mkdir -p $HOME/Go/src/github.com/user
  2) Setup your paths in your ~/.bash_profile (you can open this with a text editor and copy and paste)
    export GOPATH=$HOME/Go
    export GOROOT=/usr/local/opt/go/libexec
    export PATH=$PATH:$GOPATH/bin
    export PATH=$PATH:$GOROOT/bin
  3) Install Go (type this into your terminal)
  brew install go
```
3. Install [Nox App Player](http://www.bignox.com)
4. Make sure that there is only 1 page on the emulator and L2 Revolution is on the first page
5. in your teminal, type `cd $HOME/Go/src/github.com` and clone the repo into that directory
  to clone, paste this into your terminal: `git clone https://github.com/qhenkart/l2-bot.git`

### Configuration
1. (optional) update the conf.json file with the hour that you want the scripts to run. Options range from 0 (midnight) - 23 (11:00pm). If you neglect to set this configuration, it will default to midnight in your local time
2. (optional) Register with [PushOver](https://pushover.net/) to add push notifications to your phone. Download the application on your phone, and type in the user and token they provide into the conf.json file
3. (optional) Connect log.txt to your favorite cloud storage platform (Dropbox, icloud etc) to have live logs on your phone


### Commands
1. Midnight Cron job:
  To run the Cron job, open up your terminal and run
  `go run main.go`
  The process will run in the background until midnight occurs. It will then open up Nox App Player and run the scripts. You cannot use your computer while the scripts are running since it replicates human behavior. If you restart your computer, you must run this again.

2. Individual script:
  Utilize the --script flag. Accepted options are:
  1. *grind* : kills mobs (but makes disconnections impossible) — Can be used for any situation where you want to auto-battle and walk away without worrying about disconnections
  2. *quests* : runs the main quest line
  3. *weeklys* : runs weeklys and daily quests
  4. *sub* : runs sub-quests (you must reset counts yourself)
  5. *dungeons* : runs temple and summoning dungeon
  The process will end once the script is complete
ex. `go run main.go --script grind`


# Contributions

### Known Bugs / Contribution Requests
1. To determine if weeklys are finished or not, it looks for the [Weeklys] icon on the left sidebar. This can sometimes be hidden under the Daily Quests list which would cause Weeklys to not be run. We need to find another way to determine if there are any weeklys left

2. There are 2 or 3 methods that make use of Regions in SikuliX. Regions are less favorable than what is mostly used which is app.Window(). However in these cases, using a Region is required to avoid SikuliX confusing one image with another. This might be improved by changing the minTargetSize. Regions might cause failures on various screen sizes

3. Would be nice to have some response to getting PKed. Currently nothing happens and the script will finish. It might be nice to add a configuration option that allows people to choose Safe Mode or War mode. Safe mode could mean that the player should revive and continue the script. War mode could mean to end the script upon death, or change channels

Please make pull requests to contribute. You can fork as well but I am 100% open to feature requests and different avenues forward
