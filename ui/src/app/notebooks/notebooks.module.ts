import { NgModule } from '@angular/core';

import { SharedModule } from '@tide-shared/shared.module';
import { NotebooksService } from './notebooks.service'
import { NotebooksRoutingModule, declarations, providers } from './notebooks-routing.module';
import { NotebooksComponent } from './notebooks.component';
import { ListComponent } from './list/list.component'
@NgModule({
  declarations: [
    // ...declarations
    NotebooksComponent,
    ListComponent
  ],
  providers: [
    NotebooksService,
    ...providers,
  ],
  imports: [
    SharedModule,
    NotebooksRoutingModule
  ],
})
export class NotebooksModule {}
