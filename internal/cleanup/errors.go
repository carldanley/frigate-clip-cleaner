package cleanup

import "errors"

var ErrorFrigateConfigDoesNotExist = errors.New("frigate config path does not exist")
var ErrorFrigateClipDirDoesNotExist = errors.New("frigate clip path does not exist")
