import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { NotebooksComponent } from './notebooks.component';
import { ListComponent } from './list/list.component'
const routes: Routes = [
  {
    path: '',
    component: NotebooksComponent,
    children: [
      {
        path: '',
        component: ListComponent,
      },
    ],
  },
];

export const declarations = [
];

export const providers = [
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class NotebooksRoutingModule {}
