package esquery

import (
	"testing"
)

func TestCombinedFields(t *testing.T) {
	runMapTests(t, []mapTest{
		{
			"simple combined_fields",
			CombinedFields("value1").Fields("title"),
			map[string]interface{}{
				"combined_fields": map[string]interface{}{
					"fields": []string{"title"},
					"query":  "value1",
				},
			},
		},
		{
			"combined_fields all params",
			CombinedFields("original").
				Query("test").
				Fields("title", "body").
				AutoGenerateSynonymsPhraseQuery(true).
				Boost(6.4).
				Operator(OperatorAnd).
				MinimumShouldMatch("3<90%").
				ZeroTermsQuery(ZeroTermsAll),
			map[string]interface{}{
				"combined_fields": map[string]interface{}{
					"auto_generate_synonyms_phrase_query": true,
					"boost":                               6.4,
					"minimum_should_match":                "3<90%",
					"operator":                            "AND",
					"zero_terms_query":                    "all",
					"query":                               "test",
					"fields":                              []string{"title", "body"},
				},
			},
		},
	})
}

