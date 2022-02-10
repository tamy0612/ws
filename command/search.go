package command

import (
    "errors"
    "fmt"
    "regexp"

    "github.com/urfave/cli/v2"
    "github.com/jmoiron/sqlx"

    "github.com/tamy0612/ws/dict"
)

var (
    match_reg = regexp.MustCompile(`(?i)^[a-zA-Z_]*$`)
    include_reg = regexp.MustCompile(`[a-zA-Z]*`)
    exclude_reg = regexp.MustCompile(`[a-zA-Z]*`)
    compounds = ".-' "
)

func SearchFlags() []cli.Flag {
    return []cli.Flag{
        &cli.StringFlag{Name: "dict", Usage: "set dictionary file path"},
        &cli.StringFlag{Name: "match", Usage: "filter by matched pattern with placeholders (placeholder: '_')", Aliases: []string{"m"}},
        &cli.StringFlag{Name: "unmatch", Usage: "filter by unmatched pattern with placeholders (placeholder: '_')", Aliases: []string{"u"}},
        &cli.StringFlag{Name: "include", Usage: "filter by included characters", Aliases: []string{"i"}},
        &cli.StringFlag{Name: "exclude", Usage: "filter by non-used characters", Aliases: []string{"e"}},
        &cli.UintFlag{Name: "length", Usage: "filter by word length (ignored when `-m` used)", Aliases: []string{"l"}},
        &cli.BoolFlag{Name: "exclude-compounds", Usage: "exclude compounds", Value: false},
        &cli.BoolFlag{Name: "verbose", Usage: "displays some logs", Value: false},
    }
}

func Search(db *sqlx.DB) func(*cli.Context) error {
    return func(ctx *cli.Context) error {
        query, err := buildQuery(ctx)
        if err != nil {
            return err
        }
        candidates := []dict.Item{}
        if err := db.Select(&candidates, query); err != nil {
            return err
        }
        for _, candidate := range candidates {
            fmt.Println(candidate.Word, " :: ", candidate.Mean)
        }
        return nil
    }
}

func buildQuery(ctx *cli.Context) (string, error) {
    m := ctx.String("match")
    if !match_reg.MatchString(m) {
        return "", invalidQuery("`match` should consist of either [A-Z] or '_'")
    }
    u := ctx.String("unmatch")
    if !match_reg.MatchString(u) {
        return "", invalidQuery("`unmatch` should consist of either [A-Z] or '_'")
    }
    i := ctx.String("include")
    if !include_reg.MatchString(i) {
        return "", invalidQuery("`include` should be [A-Z]")
    }
    e := ctx.String("exclude")
    if !exclude_reg.MatchString(e) {
        return "", invalidQuery("`exclude` should be [A-Z]")
    }

    conds := []dict.QueryCondition{}
    if m != "" {
        conds = append(conds, dict.Match("word", m))
    } else if l := ctx.Uint("length"); l > 0 {
        conds = append(conds, dict.Length("word", uint32(l)))
    }
    if u != "" {
        conds = append(conds, dict.Unmatch("word", u))
    }
    if i != "" {
        conds = append(conds, dict.Include("word", i))
    }
    if e != "" {
        conds = append(conds, dict.Exclude("word", e))
    }
    if ctx.Bool("exclude-compounds") {
        conds = append(conds, dict.Exclude("word", compounds))
    }
    query := dict.Query([]string{"word", "mean"}, "items", conds...)
    if ctx.Bool("verbose") {
        fmt.Println("# Executed query: `" + query + "`")
    }
    return query, nil
}

func invalidQuery(msg string) error {
    return errors.New("Invalid query: " + msg)
}
