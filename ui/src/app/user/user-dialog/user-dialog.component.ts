import { Component, OnInit, Input, Output, EventEmitter } from '@angular/core';
import { FormBuilder, Validators, FormGroup } from '@angular/forms';
import { UserService } from '../user.service';
import { TranslateService } from '@ngx-translate/core';
import { cloudPlatform, defaultCloudPlatformURL, defaultResType, resTypes, defaultRoleType4Site, roleTypes } from '@tide-config/cloudPlatform';
import { UserListComponent } from '../user-list/user-list.component';

@Component({
  selector: 'tide-user-dialog',
  templateUrl: './user-dialog.component.html',
  styleUrls: ['./user-dialog.component.scss'],
})
export class UserDialogComponent implements OnInit {

  constructor(
    private readonly fb: FormBuilder,
    public readonly translate: TranslateService,
    public readonly userService: UserService,
    public readonly userList: UserListComponent,
  ) {
    this.userForm = this.fb.group({
      name: ['', [Validators.required]],
      org: ['', [Validators.required]],
      role: [defaultRoleType4Site, [Validators.required]],
      email: ['', [Validators.required]],
      phone: ['', [Validators.required]],
    });

    this.orgmap = userList.orgList;
    this.orgNames = Object.keys(userList.orgList);
    this.roleTypeList = Object.keys(roleTypes);
    this.roleType = roleTypes;
  }

  @Input() opened = false;
  @Output() save = new EventEmitter();
  @Output() cancel = new EventEmitter();


  userForm: FormGroup;
  orgmap: any;
  orgNames: string[];
  roleTypeList: string[];
  roleType: any;

  readonly vo = {
    serverError: '',
    spinning: false,
  };

  ngOnInit(): void {

  }

  onCancel() {
    this.close();
  }

  async onSave() {
    console.log(this.userForm);
    const { value } = this.userForm;
    this.resetModal();
    this.vo.spinning = true;
    await this.userService.addUser(value).then(() => {
      this.save.emit('');
      this.close();
      this.vo.spinning = false;
    }, (error) => {
      this.vo.serverError = error;
      this.vo.spinning = false;
    });
  }

  private close() {
    this.cancel.emit();
  }

  private resetModal() {
    this.vo.serverError = '';
    this.vo.spinning = false;
  }

}
