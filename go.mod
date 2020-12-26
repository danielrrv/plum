module github/danielrrv/plum

go 1.15

require github.com/manifoldco/promptui v0.8.0


//export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))