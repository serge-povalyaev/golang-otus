package hw10programoptimization

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
)

//easyjson:json
type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
	Address  string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	u, err := getUsers(r)
	if err != nil {
		return nil, fmt.Errorf("get users error: %w", err)
	}
	return countDomains(u, domain)
}

type users [100_000]User

func getUsers(r io.Reader) (result users, err error) {
	scanner := bufio.NewScanner(r)
	i := 0
	for scanner.Scan() {
		user := &User{}
		if err = user.UnmarshalJSON(scanner.Bytes()); err != nil {
			return
		}
		result[i] = *user
		i++
	}

	if err = scanner.Err(); err != nil {
		return
	}

	return
}

func countDomains(u users, domain string) (DomainStat, error) {
	result := make(DomainStat)
	re, err := regexp.Compile("\\." + domain)
	if err != nil {
		return nil, err
	}

	for _, user := range u {
		matched := re.Match([]byte(user.Email))
		if err != nil {
			return nil, err
		}

		if matched {
			num := result[strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])]
			num++
			result[strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])] = num
		}
	}
	return result, nil
}
