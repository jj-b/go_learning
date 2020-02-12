package todo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

/*Information about a ToDo item consists of multiple different data types (bool, time, etc) so we store these in a struct*/
//The name of this is lower case, so it can not be accessed outside the package
type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

/*Users will interact with ToDo items via the 'list' type*/
