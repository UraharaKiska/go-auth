package service

//go:generate sh -c "rm -rf mocks && mkdir -p mock"
//go:generate ../../bin/minimock -i AuthService -o ./mock -s "_minimock.go"
