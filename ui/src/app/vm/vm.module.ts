import { NgModule } from '@angular/core';

import { SharedModule } from '@tide-shared/shared.module';
import { declarations, providers, VmRoutingModule } from './vm-routing.module';


@NgModule({
  declarations: [
    ...declarations,
  ],
  providers: [
    ...providers,
  ],
  imports: [
    SharedModule,
    VmRoutingModule,
  ],
})
export class VmModule {}