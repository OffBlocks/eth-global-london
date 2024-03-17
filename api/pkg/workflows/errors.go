package workflows

import "errors"

var ErrRejected = errors.New("workflow rejected")

var ErrExpired = errors.New("workflow expired")
