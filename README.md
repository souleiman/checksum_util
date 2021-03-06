checksum_util
=========
Various minimalistic checksum for Windows user who don't want to use Cygwin.

Requirements
------------
* Go 1.4.x or greater

## Installation
    # This assumes you have GOPATH and GOBIN setup
    # Clone the project
    > git clone https://github.com/souleiman/checksum_util.git
    > cd checksum_util

    # For Windows
    ./install.bat

    # For Linux / Mac users (Note: you are better off using the checksum provided by your OS)
    chmod +x install.sh
    ./install.sh
    
## Usage
Once installed, you should be able to run the following:

|checksums|
:----------:
|md5sum|
|sha1sum|
|sha224sum|
|sha256sum|
|sha334sum|
|sha512sum|

Note: You will need to set GOBIN to your path (Environment Variables for Windows; bash profile for linux/mac)

> md5sum filename
