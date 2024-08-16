## Tool Installation Steps

1. **OpenSSH (Open Secure Shell):**

 - Installing OpenSSH on my Ubuntu machine: `sudo apt-get install openssh-client openssh-server`
 - Generating a public key through ssh-keygen agent: `ssh-keygen -t rsa -b 4096 -C "github.com@theinvincible"`
 - Printing the public key to stdout: `cat ~/.ssh/id_rsa.pub`
 - Checking if my public key has been connected remotely to GitHub:  `ssh -T git@github.com` (The response: `Hi TheInvincibleRalph! You've successfully authenticated, but GitHub does not provide shell access`indicates that my secure connection with GitHub was successful).

 -   *The `-T` option in the `ssh -T git@github.com `command is used to disable pseudo-terminal allocation.*

 >  A `pseudo-terminal` is an interface that allows the remote shell to behave as if it's interacting with a terminal (like my command line), even if it's not (well, that is why it is called `pseudo`). So, the `-T` tag is just saying *"Hey, don't bother about that pty, ain't having any interactive seesions with you, I am only running one command."*




2. **VSCode and Git:**

For VSCode, all I had to do was just navigate to the official VSCode download page at [VSCode](https://code.visualstudio.com/), then clicked on the `Download for Windows` button. Once the installer is downloaded, I then double-clicked to launch it. The rest was just to follow the installation wizard, boom my VSCode is up and running!

Git is easy, I just visited the [official Git page](https://git-scm.com/) and clicked on `Download for Windows` and waited for it to be downloaded.

**Screenshots**

> VSCode
![Screenshots](/screenshots/vscode.png)
> Git
![Screenshots](/screenshots/git%20version.png)
