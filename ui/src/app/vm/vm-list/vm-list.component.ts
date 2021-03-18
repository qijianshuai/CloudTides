import { Component, OnInit, OnDestroy} from '@angular/core';
import { TranslateService } from '@ngx-translate/core';
import { VENDOR_USAGE_REFRESH_PERIOD } from '@tide-shared/config/const';
import { Observable, of} from 'rxjs';
import { Item, VmService } from '../vm.service';

@Component({
  selector: 'tide-vm-list',
  templateUrl: './vm-list.component.html',
  styleUrls: ['./vm-list.component.scss']
})
export class VmListComponent implements OnInit {

  constructor(
    private vmService: VmService,
    public readonly translate: TranslateService,
  ) { }

  list$: Observable<Item[]> = of([]);
  opened = false;
  refreshInterval: number;

  async save() {
    await this.refreshList();
  }

  cancel() {
    this.opened = false;
  }

  async ngOnInit() {
    await this.refreshList();
  }

  async refreshList() {
    this.list$ = of(await this.vmService.getList());
    this.refreshInterval = window.setInterval(async () => {
      this.list$ = of(await this.vmService.getList());
    }, VENDOR_USAGE_REFRESH_PERIOD);
  }

  ngOnDestroy(): void {
    window.clearInterval(this.refreshInterval);
  }
}
