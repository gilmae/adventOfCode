#! zsh

#cp template $1.go
cp template2 $1.rb

sed -i '' "s/xxxxx/$1/g" $1.*
touch inputs/$1.input
touch inputs/$1.simple.input