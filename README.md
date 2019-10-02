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
| ---- | ----- | --- | ----- | ------ | ----------- |
| OP_NOP	| 97	| 0x61	| Nothing	| Nothing	| Does nothing.|
| OP_IF	| 99	| 0x63 | <expression> if [statements] [else [statements]]* endif | <expression> if [statements] [else [statements]]* endif |	If the top stack value is not False, the statements are executed. The top stack value is removed. |
| OP_NOTIF |	100	| 0x64 | <expression> notif [statements] [else [statements]]* endif	| <expression> notif [statements] [else [statements]]* endif | If the top stack value is False, the statements are executed. The top stack value is removed. |
| OP_ELSE	| 103	| 0x67 | <expression> if [statements] [else [statements]]* endif | <expression> if [statements] [else [statements]]* endif | If the preceding OP_IF or OP_NOTIF or OP_ELSE was not executed then these statements are and if the preceding OP_IF or OP_NOTIF or OP_ELSE was executed then these statements are not. |
| OP_ENDIF	| 104	| 0x68 | <expression> if [statements] [else [statements]]* endif | <expression> if [statements] [else [statements]]* endif |  Ends an if/else block. All blocks must end, or the transaction is invalid. An OP_ENDIF without OP_IF earlier is also invalid. |
| OP_VERIFY	| 105	| 0x69 |True / false |	Nothing / fail	| Marks transaction as invalid if top stack value is not true. The top stack value is removed. |
| OP_RETURN	| 106	| 0x6a | Nothing	| fail	| Marks transaction as invalid. Since bitcoin 0.9, a standard way of attaching extra data to transactions is to add a zero-value output with a scriptPubKey consisting of OP_RETURN followed by data. Such outputs are provably unspendable and specially discarded from storage in the UTXO set, reducing their cost to the network. Since 0.12, standard relay rules allow a single output with OP_RETURN, that contains any sequence of push statements (or OP_RESERVED[1]) after the OP_RETURN provided the total scriptPubKey length is at most 83 bytes. |

### Stack
| Word	| Opcode	| Hex	| Input	| Output	| Description |
| ---- | ----- | --- | ----- | ------ | ----------- |
| OP_TOALTSTACK	| 107	| 0x6b	| x1	| (alt)x1	| Puts the input onto the top of the alt stack. Removes it from the main stack. |
| OP_FROMALTSTACK	| 108	| 0x6c	| (alt)x1	| x1	| Puts the input onto the top of the main stack. Removes it from the alt stack. |
| OP_IFDUP	| 115	| 0x73	| x	| x / x x	| If the top stack value is not 0, duplicate it. |
| OP_DEPTH	| 116	| 0x74	| Nothing	| <Stack size> |	Puts the number of stack items onto the stack. |
| OP_DROP	| 117	| 0x75	| x	| Nothing	|Removes the top stack item. |
| OP_DUP	| 118	| 0x76	| x	| x x	| Duplicates the top stack item. |
| OP_NIP	| 119	| 0x77	| x1 x2	| x2	| Removes the second-to-top stack item. |
| OP_OVER	| 120	| 0x78	| x1 x2	| x1 x2 x1	| Copies the second-to-top stack item to the top. |
| OP_PICK	| 121	| 0x79	| xn ... x2 x1 x0 <n> |	xn ... x2 x1 x0 xn	| The item n back in the stack is copied to the top. |
| OP_ROLL	| 122	| 0x7a	| xn ... x2 x1 x0 <n>	| ... x2 x1 x0 xn	| The item n back in the stack is moved to the top. |
| OP_ROT	| 123	| 0x7b	| x1 x2 x3	| x2 x3 x1	| The top three items on the stack are rotated to the left. |
| OP_SWAP	| 124	| 0x7c	| x1 x2	| x2 x1	| The top two items on the stack are swapped. |
| OP_TUCK	| 125	| 0x7d	| x1 x2	| x2 x1 x2	| The item at the top of the stack is copied and inserted before the second-to-top item. |
| OP_2DROP	| 109	| 0x6d	| x1 x2	| Nothing	| Removes the top two stack items. |
| OP_2DUP	| 110	| 0x6e	| x1 x2	| x1 x2 x1 x2	| Duplicates the top two stack items. |
| OP_3DUP	| 111	| 0x6f	| x1 x2 x3	| x1 x2 x3 x1 x2 x3	| Duplicates the top three stack items. |
| OP_2OVER	| 112	| 0x70	| x1 x2 x3 x4	| x1 x2 x3 x4 x1 x2	| Copies the pair of items two spaces back in the stack to the front. |
| OP_2ROT	| 113	| 0x71	| x1 x2 x3 x4 x5 x6	| x3 x4 x5 x6 x1 x2	| The fifth and sixth items back are moved to the top of the stack. |
| OP_2SWAP	| 114	| 0x72	| x1 x2 x3 x4	| x3 x4 x1 x2	| Swaps the top two pairs of items. |
