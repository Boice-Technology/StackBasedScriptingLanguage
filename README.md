# StackBasedScriptingLanguage
Implementation of stack based scripting language for validation of transactions(clearing)


### Constants
| Word |Opcode | Hex | Input | Output | Description |
| ---- | ----- | --- | ----- | ------ | ----------- |
| OP_0, OP_FALSE |	0 | 0x00 |	Nothing.	| (empty value)	| An empty array of bytes is pushed onto the stack. (This is not a no-op: an item is added to the stack.) |
| N/A |	1-75 |	0x01-0x4b	| (special)	| data	| The next opcode bytes is data to be pushed onto the stack |
| OP_PUSHDATA1 |	76	| 0x4c	| (special)	| data	| The next byte contains the number of bytes to be pushed onto the stack. |
| OP_PUSHDATA2	| 77	| 0x4d	| (special)	| data	| The next two bytes contain the number of bytes to be pushed onto the stack in little endian order. |
| OP_PUSHDATA4	| 78	| 0x4e	| (special)	| data	| The next four bytes contain the number of bytes to be pushed onto the stack in little endian order. |
| OP_1NEGATE	| 79	| 0x4f	| Nothing. |	-1	| The number -1 is pushed onto the stack. |
| OP_1, OP_TRUE	81	| 0x51	| Nothing.	| 1	| The number 1 is pushed onto the stack. |
| OP_2-OP_16	| 82-96	| 0x52-0x60	| Nothing.	| 2-16	| The number in the word name (2-16) is pushed onto the stack. |

### Flow control
| Word	| Opcode	| Hex	| Input	| Output	| Description |
| ---- | ----- | --- | ----- | ------ | ----------- 
| OP_NOP	| 97	| 0x61	| Nothing	| Nothing	| Does nothing.|
| OP_IF	| 99	| 0x63 | <expression> if [statements] [else [statements]]* endif | <expression> if [statements] [else [statements]]* endif |	If the top stack value is not False, the statements are executed. The top stack value is removed. |
| OP_NOTIF |	100	| 0x64 | <expression> notif [statements] [else [statements]]* endif	| <expression> notif [statements] [else [statements]]* endif | If the top stack value is False, the statements are executed. The top stack value is removed. |
| OP_ELSE	| 103	| 0x67 | <expression> if [statements] [else [statements]]* endif | <expression> if [statements] [else [statements]]* endif | If the preceding OP_IF or OP_NOTIF or OP_ELSE was not executed then these statements are and if the preceding OP_IF or OP_NOTIF or OP_ELSE was executed then these statements are not. |
| OP_ENDIF	| 104	| 0x68 | <expression> if [statements] [else [statements]]* endif | <expression> if [statements] [else [statements]]* endif |  Ends an if/else block. All blocks must end, or the transaction is invalid. An OP_ENDIF without OP_IF earlier is also invalid. |
| OP_VERIFY	| 105	| 0x69 |True / false |	Nothing / fail	| Marks transaction as invalid if top stack value is not true. The top stack value is removed. |
| OP_RETURN	| 106	| 0x6a | Nothing	| fail	| Marks transaction as invalid. Since bitcoin 0.9, a standard way of attaching extra data to transactions is to add a zero-value output with a scriptPubKey consisting of OP_RETURN followed by data. Such outputs are provably unspendable and specially discarded from storage in the UTXO set, reducing their cost to the network. Since 0.12, standard relay rules allow a single output with OP_RETURN, that contains any sequence of push statements (or OP_RESERVED[1]) after the OP_RETURN provided the total scriptPubKey length is at most 83 bytes. |
