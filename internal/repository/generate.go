package repository

//go:generate sh -c "rm -rf mocks && mkdir -p mock"
//go:generate ../../bin/minimock -i UserRepository -o ./mock -s "_minimock.go"
