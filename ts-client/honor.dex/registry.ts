import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreatePool } from "./types/honor/dex/tx";
import { MsgUpdatePool } from "./types/honor/dex/tx";
import { MsgDeletePool } from "./types/honor/dex/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/honor.dex.MsgCreatePool", MsgCreatePool],
    ["/honor.dex.MsgUpdatePool", MsgUpdatePool],
    ["/honor.dex.MsgDeletePool", MsgDeletePool],
    
];

export { msgTypes }