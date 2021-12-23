import { RoomsInterface } from "./IRoom";
import { EquipmentsInterface } from "./IEquipment";


export interface RoomEquipmentsInterface {
    ID: number,
    RoomID: number,
    Room: RoomsInterface,
    EquipmentID: number,
    Equipment: EquipmentsInterface,
  }
  