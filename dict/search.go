package dict

import (
    "fmt"
    "strings"
)

type QueryCondition func() string

func Match(label string, pattern string) QueryCondition {
    return func() string {
        return fmt.Sprintf(`%s LIKE "%s"`, label, pattern)
    }
}

func Include(label string, chars string) QueryCondition {
    return func() string {
        conds := []string{}
        for _, c := range chars {
            conds = append(conds, fmt.Sprintf(`%s LIKE "%%%s%%"`, label, string(c)))
        }
        return strings.Join(conds, " AND ")
    }
}

func Exclude(label string, chars string) QueryCondition {
    return func() string {
        conds := []string{}
        for _, c := range chars {
            conds = append(conds, fmt.Sprintf(`%s NOT LIKE "%%%s%%"`, label, string(c)))
        }
        return strings.Join(conds, " AND ")
    }
}

func Length(label string, length uint32) QueryCondition {
    return func() string {
        return fmt.Sprintf("length(%s) = %d", label, length)
    }
}

func Query(labels []string, table string, conditions ...QueryCondition) string {
    query := `SELECT ` + strings.Join(labels, ",") + " FROM " + table
    conds := []string{` WHERE 1 = 1`}
    for _, cond := range conditions {
        conds = append(conds, cond())
    }
    return query + strings.Join(conds, " AND ")
}
