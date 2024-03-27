# slug_generator
Generate a slug of any length of words, implemented in Go.

## Installation
```bash
go get github.com/jairad26/slug-generator
```

## Usage
```go
package main

import (
    "fmt"
    sg "github.com/jairad26/slug_generator"
)

func main() {
    seed := "literally anything, timestamp, etc."
    generator := sg.NewSlugGenerator(seed, sg.ADJECTIVES, sg.NOUNS)
    slug := generator.GenerateSlug(5)

    fmt.Println(slug)
}
```

## Slug Logic
1. A slug must end with a noun.
2. You can have at most 2 adjectives describing a noun.
3. A noun is followed by "of" (unless it is the end of the slug).
4. A slug must begin with either a noun or an adjective.
5. "of" can not be followed by "of".

Some quirks arise from these, feel free to open an issue if you find any.

## Example Slug
Number of Words | Example Slug
----------------|--------------
1               | `tree`
2               | `dead-tree`
3               | `misty-dead-tree` OR `tree-of-thunder`
4               | `misty-tree-of-thunder` OR `tree-of-misty-thunder`
5               | `misty-tree-of-ancient-thunder` OR `tree-of-misty-ancient-thunder` OR `ancient-misty-tree-of-thunder`

## License
I have none. Do whatever you want with this code.


## Contributing
Feel free to contribute to this project, I just made it up at 1 am one morning, so there are definitely improvements to be made. Open a PR and I'll review it as soon as possible. Or open an issue if you find a bug or have a feature request.

## Author
Jai Radhakrishnan

