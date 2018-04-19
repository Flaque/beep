// Generated from /home/flaque/projects/go/src/github.com/Flaque/boop/parser/BeepBoop.g4 by ANTLR 4.7.1
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class BeepBoopLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.7.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, COMMENT=2, NEWLINE=3, WHITESPACE=4, IF=5, DO=6, END=7, FUNC=8, 
		RETURN=9, FOR=10, IS=11, PLUS=12, MINUS=13, MULT=14, DIVIDE=15, ASSIGN=16, 
		PIPE=17, LPAREN=18, RPAREN=19, STRING=20, INT=21;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	public static final String[] ruleNames = {
		"T__0", "COMMENT", "NEWLINE", "WHITESPACE", "IF", "DO", "END", "FUNC", 
		"RETURN", "FOR", "IS", "PLUS", "MINUS", "MULT", "DIVIDE", "ASSIGN", "PIPE", 
		"LPAREN", "RPAREN", "STRING", "INT"
	};

	private static final String[] _LITERAL_NAMES = {
		null, "'$'", null, null, null, "'if'", "'do'", "'end'", "'func'", "'return'", 
		"'for'", "'is'", "'+'", "'-'", "'*'", "'/'", "'='", "'|'", "'('", "')'"
	};
	private static final String[] _SYMBOLIC_NAMES = {
		null, null, "COMMENT", "NEWLINE", "WHITESPACE", "IF", "DO", "END", "FUNC", 
		"RETURN", "FOR", "IS", "PLUS", "MINUS", "MULT", "DIVIDE", "ASSIGN", "PIPE", 
		"LPAREN", "RPAREN", "STRING", "INT"
	};
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}


	public BeepBoopLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "BeepBoop.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\27~\b\1\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\3\2\3\2\3\3\3\3\7\3\62\n\3\f"+
		"\3\16\3\65\13\3\3\3\3\3\3\4\3\4\6\4;\n\4\r\4\16\4<\5\4?\n\4\3\5\6\5B\n"+
		"\5\r\5\16\5C\3\5\3\5\3\6\3\6\3\6\3\7\3\7\3\7\3\b\3\b\3\b\3\b\3\t\3\t\3"+
		"\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\f\3\f\3\f"+
		"\3\r\3\r\3\16\3\16\3\17\3\17\3\20\3\20\3\21\3\21\3\22\3\22\3\23\3\23\3"+
		"\24\3\24\3\25\6\25v\n\25\r\25\16\25w\3\26\6\26{\n\26\r\26\16\26|\2\2\27"+
		"\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20"+
		"\37\21!\22#\23%\24\'\25)\26+\27\3\2\6\4\2\f\f\17\17\4\2\13\13\"\"\4\2"+
		"C\\c|\3\2\62;\2\u0083\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2"+
		"\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25"+
		"\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2"+
		"\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2"+
		"\2\3-\3\2\2\2\5/\3\2\2\2\7>\3\2\2\2\tA\3\2\2\2\13G\3\2\2\2\rJ\3\2\2\2"+
		"\17M\3\2\2\2\21Q\3\2\2\2\23V\3\2\2\2\25]\3\2\2\2\27a\3\2\2\2\31d\3\2\2"+
		"\2\33f\3\2\2\2\35h\3\2\2\2\37j\3\2\2\2!l\3\2\2\2#n\3\2\2\2%p\3\2\2\2\'"+
		"r\3\2\2\2)u\3\2\2\2+z\3\2\2\2-.\7&\2\2.\4\3\2\2\2/\63\7%\2\2\60\62\n\2"+
		"\2\2\61\60\3\2\2\2\62\65\3\2\2\2\63\61\3\2\2\2\63\64\3\2\2\2\64\66\3\2"+
		"\2\2\65\63\3\2\2\2\66\67\b\3\2\2\67\6\3\2\2\28?\t\2\2\29;\t\2\2\2:9\3"+
		"\2\2\2;<\3\2\2\2<:\3\2\2\2<=\3\2\2\2=?\3\2\2\2>8\3\2\2\2>:\3\2\2\2?\b"+
		"\3\2\2\2@B\t\3\2\2A@\3\2\2\2BC\3\2\2\2CA\3\2\2\2CD\3\2\2\2DE\3\2\2\2E"+
		"F\b\5\3\2F\n\3\2\2\2GH\7k\2\2HI\7h\2\2I\f\3\2\2\2JK\7f\2\2KL\7q\2\2L\16"+
		"\3\2\2\2MN\7g\2\2NO\7p\2\2OP\7f\2\2P\20\3\2\2\2QR\7h\2\2RS\7w\2\2ST\7"+
		"p\2\2TU\7e\2\2U\22\3\2\2\2VW\7t\2\2WX\7g\2\2XY\7v\2\2YZ\7w\2\2Z[\7t\2"+
		"\2[\\\7p\2\2\\\24\3\2\2\2]^\7h\2\2^_\7q\2\2_`\7t\2\2`\26\3\2\2\2ab\7k"+
		"\2\2bc\7u\2\2c\30\3\2\2\2de\7-\2\2e\32\3\2\2\2fg\7/\2\2g\34\3\2\2\2hi"+
		"\7,\2\2i\36\3\2\2\2jk\7\61\2\2k \3\2\2\2lm\7?\2\2m\"\3\2\2\2no\7~\2\2"+
		"o$\3\2\2\2pq\7*\2\2q&\3\2\2\2rs\7+\2\2s(\3\2\2\2tv\t\4\2\2ut\3\2\2\2v"+
		"w\3\2\2\2wu\3\2\2\2wx\3\2\2\2x*\3\2\2\2y{\t\5\2\2zy\3\2\2\2{|\3\2\2\2"+
		"|z\3\2\2\2|}\3\2\2\2},\3\2\2\2\t\2\63<>Cw|\4\b\2\2\2\3\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}