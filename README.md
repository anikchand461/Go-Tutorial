# GO TUTORIAL

_**If you want to contribute in this repository , Go to the contribute branch**_

Go is a multi-purpose and general programming language , developed by Google in 2007. ![App Screenshot](https://cdn1.iconfinder.com/data/icons/google-s-logo/150/Google_Icons-09-20.png)

Nowadays , it is majorly used to build robust **backends** 💿 of different applications which needs features like **concurrency** for its scalability (to provide services to huge traffic). Go was built to act as a modern replacement of C by Google but it got its own place 🥈 (yet we know , why C is the king of all languages) 👑.

Applications, which uses technologies like **Blockchain , Real-time communication or even ML (use for generic purposes also 😄 )** , can take advantages of Go.

You can run **go** in any OS (Windows , Linux , Mac) 💻. Here , I will describe the **installation steps in Linux** only . For , other ways , have a look in the url attached - [link](https://go.dev/doc/install) 👀

## Installation in linux (boot or WSL)
The steps are as follows -

### Install the Script
- Go to ths link ➡️ [link](https://go.dev/dl/) . Click on the **.tar.gz** file under **Linux** section and one tar file will be downloaded in your system.

  ![image](https://github.com/user-attachments/assets/254fb2f1-5220-4f99-8f04-0b81694f0916)

- For **WSL** users , check where the tar file is downloaded. Move the file to the location where your linux distro is saved.
- Then run the following command -
```bash
  tar -xzf <file-name>.tar.gz
```
- For **Fedora OS** users , run the following command -
```bash
  sudo dnf install golang
```

### Configure GOROOT and GOPATH

- **GOROOT** is the location where your extracted go executables are placed (extracted from **tar.gz** file)

To find the GOROOT location , run this command -
```bash
  whereis go
```
- Now let's configure the paths. Go projects typically follow a folder structure which is divided into 3 sections - **bin** , **pkg** , **src**.

- - **src** - contains actual go SOURCE CODES .

- - **pkg** - compiled versions of our imported go PACKAGES .

- - **bin** - compiled executable binaries of of our go SCRIPTS .

- In first **GOROOT** , we specify the location where the executable binary packages of go resides , typically ```/usr/local/go``` .

- Then we add the location (GOROOT) to the linux's **PATH** variable , to let it know the location .

- The **GOPATH** denotes the location of our go projects and created packages.

- open your **.bashrc** ( and **.zshrc** if you use zsh file-system or any other additional file-system ) and paste these codes ( SUBSTITUTE THE FOLDER LOCATIONS ONLY ❗ )

```bash

  export GOROOT=/path/to/golang
  export PATH=$PATH:$GOROOT/bin
  
  export GOPATH=$GOPATH:/path/to/go-folder/programs

```

### Download essential libraries -

Open **/src** of your go workspace . open terminal . run the following commands which will help to install basic and necessary go tools (libraries)

```bash
  go install golang.org/x/tools/cmd/goimports@latest

  go install golang.org/x/lint/golint@latest

```

### IDE
For me , Jetbrain's GOLAND is best environment where you can code in GO.
