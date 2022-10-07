import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreatePool } from "./types/dex/tx";
import { MsgUpdatePool } from "./types/dex/tx";
import { MsgDeletePool } from "./types/dex/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/meow.dex.MsgCreatePool", MsgCreatePool],
    ["/meow.dex.MsgUpdatePool", MsgUpdatePool],
    ["/meow.dex.MsgDeletePool", MsgDeletePool],
    
];

export { msgTypes }