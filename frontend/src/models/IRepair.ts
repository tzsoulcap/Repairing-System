import { RoomEquipmentsInterface } from "./IRoomEquipment";
import { RentalsInterface } from "./IRental";
import { TenantsInterface } from "./ITenant";

export interface RepairInterface {
    ID: number,
    AddedTime: Date,
    Note: string,
    TenantID: number,
    Tenant: TenantsInterface,
    RentalID: number,
    Rental: RentalsInterface
    RoomEquipmentID: number,
    RoomEquipment: RoomEquipmentsInterface,
  }
  