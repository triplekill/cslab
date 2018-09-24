brew install valgrind # install for mac https://www.gungorbudak.com/blog/2018/04/28/how-to-install-valgrind-on-macos-high-sierra/
sudo apt-get install make gcc 

make && ./hello < t.txt

# profile
valgrind --tool=memcheck ./hello
valgrind --tool=cachegrind ./hello
valgrind --tool=callgrind ./hello
