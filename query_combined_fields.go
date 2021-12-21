package esquery

import "github.com/fatih/structs"

type CombinedFieldsQuery struct {
	params combinedFieldsParams
}

// Map returns a map representation of the query; implementing the
// Mappable interface.
func (q *CombinedFieldsQuery) Map() map[string]interface{} {
	return map[string]interface{}{
		"combined_fields": structs.Map(q.params),
	}
}

type combinedFieldsParams struct {
	Qry          interface{}    `structs:"query"`
	Fields       []string       `structs:"fields"`
	Boost        float32        `structs:"boost,omitempty"`
	AutoGenerate *bool          `structs:"auto_generate_synonyms_phrase_query,omitempty"`
	Op           MatchOperator  `structs:"operator,string,omitempty"`
	MinMatch     string         `structs:"minimum_should_match,omitempty"`
	ZeroTerms    ZeroTerms      `structs:"zero_terms_query,string,omitempty"`
}

// CombinedFields creates a new query of type "combined_fields"
func CombinedFields(simpleQuery interface{}) *CombinedFieldsQuery {
	return newCombinedFields(simpleQuery)
}

func newCombinedFields(simpleQuery interface{}) *CombinedFieldsQuery {
	return &CombinedFieldsQuery{
		params: combinedFieldsParams{
			Qry: simpleQuery,
		},
	}
}

// Query sets the data to find in the query's field (it is the "query" component
// of the query).
func (q *CombinedFieldsQuery) Query(data interface{}) *CombinedFieldsQuery {
	q.params.Qry = data
	return q
}

// Fields sets the fields used in the query
func (q *CombinedFieldsQuery) Fields(a ...string) *CombinedFieldsQuery {
	q.params.Fields = append(q.params.Fields, a...)
	return q
}

// AutoGenerateSynonymsPhraseQuery sets the "auto_generate_synonyms_phrase_query"
// boolean.
func (q *CombinedFieldsQuery) AutoGenerateSynonymsPhraseQuery(b bool) *CombinedFieldsQuery {
	q.params.AutoGenerate = &b
	return q
}

// Boost
func (q *CombinedFieldsQuery) Boost(l float32) *CombinedFieldsQuery {
	q.params.Boost = l
	return q
}

// Operator sets the boolean logic used to interpret text in the query value.
func (q *CombinedFieldsQuery) Operator(op MatchOperator) *CombinedFieldsQuery {
	q.params.Op = op
	return q
}

// MinimumShouldMatch sets the minimum number of clauses that must match for a
// document to be returned.
func (q *CombinedFieldsQuery) MinimumShouldMatch(s string) *CombinedFieldsQuery {
	q.params.MinMatch = s
	return q
}

// ZeroTermsQuery sets the "zero_terms_query" option to use. This indicates
// whether no documents are returned if the analyzer removes all tokens, such as
// when using a stop filter.
func (q *CombinedFieldsQuery) ZeroTermsQuery(s ZeroTerms) *CombinedFieldsQuery {
	q.params.ZeroTerms = s
	return q
}

