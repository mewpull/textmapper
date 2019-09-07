// generated by Textmapper; DO NOT EDIT

package json

import (
	"fmt"
)

// Parser is a table-driven LALR parser for json.
type Parser struct {
	listener Listener

	next symbol
}

type SyntaxError struct {
	Line      int
	Offset    int
	Endoffset int
}

func (e SyntaxError) Error() string {
	return fmt.Sprintf("syntax error at line %v", e.Line)
}

type symbol struct {
	symbol    int32
	offset    int
	endoffset int
}

type stackEntry struct {
	sym   symbol
	state int8
	value interface{}
}

func (p *Parser) Init(l Listener) {
	p.listener = l
}

const (
	startStackSize       = 256
	startTokenBufferSize = 16
	noToken              = int32(UNAVAILABLE)
	eoiToken             = int32(EOI)
	debugSyntax          = false
)

func (p *Parser) Parse(lexer *Lexer) error {
	return p.parse(1, 44, lexer)
}

func (p *Parser) parse(start, end int8, lexer *Lexer) error {
	ignoredTokens := make([]symbol, 0, startTokenBufferSize) // to be reported with the next shift
	state := start

	var alloc [startStackSize]stackEntry
	stack := append(alloc[:0], stackEntry{state: state})
	ignoredTokens = p.fetchNext(lexer, stack, ignoredTokens)

	for state != end {
		action := tmAction[state]
		if action < -2 {
			// Lookahead is needed.
			if p.next.symbol == noToken {
				ignoredTokens = p.fetchNext(lexer, stack, ignoredTokens)
			}
			action = lalr(action, p.next.symbol)
		}

		if action >= 0 {
			// Reduce.
			rule := action
			ln := int(tmRuleLen[rule])

			var entry stackEntry
			entry.sym.symbol = tmRuleSymbol[rule]
			rhs := stack[len(stack)-ln:]
			stack = stack[:len(stack)-ln]
			if ln == 0 {
				if p.next.symbol == noToken {
					ignoredTokens = p.fetchNext(lexer, stack, ignoredTokens)
				}
				entry.sym.offset, _ = lexer.Pos()
				entry.sym.endoffset = entry.sym.offset
			} else {
				entry.sym.offset = rhs[0].sym.offset
				entry.sym.endoffset = rhs[ln-1].sym.endoffset
			}
			if err := p.applyRule(rule, &entry, rhs, lexer); err != nil {
				return err
			}
			if debugSyntax {
				fmt.Printf("reduced to: %v\n", Symbol(entry.sym.symbol))
			}
			state = gotoState(stack[len(stack)-1].state, entry.sym.symbol)
			entry.state = state
			stack = append(stack, entry)

		} else if action == -1 {
			// Shift.
			if p.next.symbol == noToken {
				p.fetchNext(lexer, stack, nil)
			}
			state = gotoState(state, p.next.symbol)
			stack = append(stack, stackEntry{
				sym:   p.next,
				state: state,
				value: lexer.Value(),
			})
			if debugSyntax {
				fmt.Printf("shift: %v (%s)\n", Symbol(p.next.symbol), lexer.Text())
			}
			if len(ignoredTokens) > 0 {
				for _, tok := range ignoredTokens {
					p.reportIgnoredToken(tok)
				}
				ignoredTokens = ignoredTokens[:0]
			}
			switch Token(p.next.symbol) {
			case JSONSTRING:
				p.listener(JsonString, p.next.offset, p.next.endoffset)
			}
			if state != -1 && p.next.symbol != eoiToken {
				p.next.symbol = noToken
			}
		}

		if action == -2 || state == -1 {
			break
		}
	}

	if state != end {
		offset, endoffset := lexer.Pos()
		err := SyntaxError{
			Line:      lexer.Line(),
			Offset:    offset,
			Endoffset: endoffset,
		}
		return err
	}

	return nil
}

func lalr(action, next int32) int32 {
	a := -action - 3
	for ; tmLalr[a] >= 0; a += 2 {
		if tmLalr[a] == next {
			break
		}
	}
	return tmLalr[a+1]
}

func gotoState(state int8, symbol int32) int8 {
	min := tmGoto[symbol]
	max := tmGoto[symbol+1]

	if max-min < 32 {
		for i := min; i < max; i += 2 {
			if tmFromTo[i] == state {
				return tmFromTo[i+1]
			}
		}
	} else {
		for min < max {
			e := (min + max) >> 1 &^ int32(1)
			i := tmFromTo[e]
			if i == state {
				return tmFromTo[e+1]
			} else if i < state {
				min = e + 2
			} else {
				max = e
			}
		}
	}
	return -1
}

func (p *Parser) fetchNext(lexer *Lexer, stack []stackEntry, ignoredTokens []symbol) []symbol {
restart:
	token := lexer.Next()
	switch token {
	case MULTILINECOMMENT, INVALID_TOKEN:
		s, e := lexer.Pos()
		tok := symbol{int32(token), s, e}
		if ignoredTokens == nil {
			p.reportIgnoredToken(tok)
		} else {
			ignoredTokens = append(ignoredTokens, tok)
		}
		goto restart
	}
	p.next.symbol = int32(token)
	p.next.offset, p.next.endoffset = lexer.Pos()
	return ignoredTokens
}

func lookaheadNext(lexer *Lexer) int32 {
restart:
	tok := lexer.Next()
	switch tok {
	case MULTILINECOMMENT, INVALID_TOKEN:
		goto restart
	}
	return int32(tok)
}

func AtEmptyObject(lexer *Lexer, next int32) bool {
	return lookahead(lexer, next, 0, 42)
}

func lookahead(l *Lexer, next int32, start, end int8) bool {
	var lexer Lexer = *l

	var allocated [64]stackEntry
	state := start
	stack := append(allocated[:0], stackEntry{state: state})

	for state != end {
		action := tmAction[state]
		if action < -2 {
			// Lookahead is needed.
			if next == noToken {
				next = lookaheadNext(&lexer)
			}
			action = lalr(action, next)
		}

		if action >= 0 {
			// Reduce.
			rule := action
			ln := int(tmRuleLen[rule])

			var entry stackEntry
			entry.sym.symbol = tmRuleSymbol[rule]
			stack = stack[:len(stack)-ln]
			state = gotoState(stack[len(stack)-1].state, entry.sym.symbol)
			entry.state = state
			stack = append(stack, entry)

		} else if action == -1 {
			// Shift.
			if next == noToken {
				next = lookaheadNext(&lexer)
			}
			state = gotoState(state, next)
			stack = append(stack, stackEntry{
				sym:   symbol{symbol: next},
				state: state,
			})
			if state != -1 && next != eoiToken {
				next = noToken
			}
		}

		if action == -2 || state == -1 {
			break
		}
	}

	return state == end
}

func (p *Parser) applyRule(rule int32, lhs *stackEntry, rhs []stackEntry, lexer *Lexer) (err error) {
	switch rule {
	case 32:
		if AtEmptyObject(lexer, p.next.symbol) {
			lhs.sym.symbol = 23 /* lookahead_EmptyObject */
		} else {
			lhs.sym.symbol = 25 /* lookahead_notEmptyObject */
		}
		return
	}
	nt := ruleNodeType[rule]
	if nt != 0 {
		p.listener(nt, lhs.sym.offset, lhs.sym.endoffset)
	}
	return
}

func (p *Parser) reportIgnoredToken(tok symbol) {
	var t NodeType
	switch Token(tok.symbol) {
	case MULTILINECOMMENT:
		t = MultiLineComment
	case INVALID_TOKEN:
		t = InvalidToken
	default:
		return
	}
	p.listener(t, tok.offset, tok.endoffset)
}
