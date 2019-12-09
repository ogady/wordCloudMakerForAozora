package morphoAnalyzer

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bluele/mecab-golang"
)

func ParseToNode(text string) (map[string]int, error) {

	wordMap := make(map[string]int)

	m, err := mecab.New("-Owakati")
	if err != nil {
		err = fmt.Errorf("MeCabの初期化（分かち書き出力モード）に失敗しました。\n %w", err)
		return wordMap, err
	}
	defer m.Destroy()

	tg, err := m.NewTagger()
	if err != nil {
		return wordMap, err
	}

	defer tg.Destroy()

	lt, err := m.NewLattice(text)
	if err != nil {
		return wordMap, err
	}

	defer lt.Destroy()

	node := tg.ParseToNode(lt)
	for {
		features := strings.Split(node.Feature(), ",")
		if features[0] == "名詞" {
			if !contains(stopWordJPN, node.Surface()) {
				// mapのキーに単語を設定して、バリューに単語のカウントを設定し、キーに対してカウントしていく
				wordMap[node.Surface()]++
			}
		}
		if node.Next() != nil {
			break
		}
	}
	return wordMap, nil
}

func contains(sl []string, s string) bool {

	for _, v := range sl {
		if s == v {
			return true
		}
	}
	return false
}

type EntryMap struct {
	name  string
	value int
}
type List []EntryMap

func GetMaxCount(m map[string]int) string {
	a := List{}

	for k, v := range m {
		e := EntryMap{k, v}
		a = append(a, e)
	}

	sort.SliceStable(a, func(i, j int) bool { return a[i].value < a[j].value })

	return a[len(a)-1].name
}
