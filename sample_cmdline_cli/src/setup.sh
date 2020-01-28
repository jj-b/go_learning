#Sets up the environment
echo export GOPATH="/users/jbottom/dev/go/sample_cmdline_cli" >> ~/.bash_profile
echo export PATH="$PATH:$GOPATH/bin" >> ~/.bash_profile
source ~/.bash_profile
#Get packages we need
cd ..
go get golang.org/x/tools/cmd/goimports
go get github.com/ogier/pflag


