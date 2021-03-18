import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { VmComponent } from './vm.component';
import { VmListComponent } from './vm-list/vm-list.component';
import { VmService } from './vm.service';
//import { TemplateCardComponent } from './template-card/template-card.component';

const routes: Routes = [
  {
    path: '',
    component: VmComponent,
    children: [
      {
        path: '',
        component: VmListComponent,
      },
    ],
  },
];

export const declarations = [
  VmComponent,
  VmListComponent,
];

export const providers = [
  VmService,
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class VmRoutingModule {}