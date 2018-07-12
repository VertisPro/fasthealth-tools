package textutils

import (
	"strings"
)

// SpellChecker provides all functions to check spellings
type SpellChecker struct {
	//E is where we keep our enchant instance
	E *Enchant
}

//NewSpellChecker starts a new instance of our spell Checker, can optionally load a dictionary file
func NewSpellChecker(dicfile string) (*SpellChecker, error) {
	e := InitEnchant()
	err := e.DictPWLLoad(dicfile)
	if err != nil {
		return nil, err
	}
	return &SpellChecker{e}, nil
}

// CheckSentence checks a sentence supplied and provides suggestions
func (S *SpellChecker) CheckSentence(sentence string) (feedback []string, err error) {

	words := strings.Fields(sentence)

	for _, word := range words {
		found, err := S.E.DictCheck(word)
		if err != nil {
			return nil, err
		}

		if found {
			feedback = append(feedback, word)
		} else {
			suggestions, err := S.E.DictSuggest(word)
			if err != nil {
				return nil, err
			}
			spellingfeedback := "(" + strings.Join(suggestions, "/") + ")"
			feedback = append(feedback, word, spellingfeedback)
		}
	}
	return feedback, nil
}
