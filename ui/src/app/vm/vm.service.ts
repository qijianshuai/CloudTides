import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { environment } from "@tide-environments/environment";
import { VM_PATH } from "@tide-shared/config/path";
import { map } from "rxjs/operators";
import { LoginService } from "../login/login.service";

@Injectable()
export class VmService{
    
    constructor(
        private readonly http: HttpClient,
        private readonly loginService: LoginService,
    ){}

    private prefix = `${environment.apiPrefix}/computeResource`;
    
    async getList() {
        const list = await this.http.get<Item[]>(environment.apiPrefix + VM_PATH, {
          headers: {
            Authorization: `Bearer ${this.loginService.token}`,
          },
        }).toPromise();
        const vms: Item[] = [];
        for (const vm of list) {
          const vmItem: Item = {
            id: vm.id,
            name: vm.name,
            ipAddress: vm.ipAddress,
            tempId: vm.tempId,
            cpuNum: vm.cpuNum,
            datacenter: vm.datacenter,
          };
          vms.push(vmItem);
        }
        return vms;
    }
    
    addItem(payload: ItemPayload) {
        const body = {
          ...payload,
        };
        return this.http.post<any>(environment.apiPrefix + VM_PATH, body, {
          headers: {
            Authorization: `Bearer ${this.loginService.token}`,
          },
        }).toPromise().then(() => {
          return Promise.resolve();
        }, (errResp) => {
          return Promise.reject(`HTTP ${errResp.status}: ${errResp.error.message}`);
        });
    }
    
    editItem(id: string, payload: ItemPayload) {
        return this.http.put<ItemDTO>(`${this.prefix}/${id}`, payload).pipe(
          map(mapItem),
        );
    }
    
      async removeItem(id: string) {
        await this.http.delete<any>(environment.apiPrefix + `/resource/vcd/` + id, {
          headers: {
            Authorization: `Bearer ${this.loginService.token}`,
        }, }).toPromise().then(
          () => {
            return Promise.resolve();
          }, (errResp) => {
            return Promise.reject(`${errResp.message}`);
          },
        );
      }
}

interface ItemDTO {
    id: number;
    name: string;
    ipAddress: string;
    tempId: string;
    cpuNum: number;
    datacenter: string;
}


function mapList(raw: ItemDTO[]): Item[] {
    return raw.map(mapItem);
}
  
function mapItem(raw: ItemDTO): Item {
    return raw;
}

export interface ItemPayload {
    name: string;
    version: string;
    vendorType: string;
    url: string;
}

export type Item = ItemDTO;