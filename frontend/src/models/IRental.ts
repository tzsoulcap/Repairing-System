import { TenantsInterface } from "./ITenant";
import { RoomsInterface } from "./IRoom";

export interface RentalsInterface {
    ID: number,
    TenantID: number,
    Tenant: TenantsInterface,
    RoomID: number,
    Room: RoomsInterface,
    State: string,
    Daytime: Date,
    Checkin: Date,
  }
  