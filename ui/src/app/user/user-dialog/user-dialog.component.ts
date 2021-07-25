import { Component, OnInit, Input, Output, EventEmitter } from "@angular/core";
import { FormBuilder, Validators, FormGroup } from "@angular/forms";
import { UserService } from "../user.service";
import { TranslateService } from "@ngx-translate/core";
import { defaultRoleType4Site, defaultRoleType4Org, roleTypes, roleTypes4Org } from "@tide-config/cloudPlatform";
import { UserListComponent } from "../user-list/user-list.component";
import { LoginService } from "src/app/login/login.service";

@Component({
  selector: "tide-user-dialog",
  templateUrl: "./user-dialog.component.html",
  styleUrls: ["./user-dialog.component.scss"],
})
export class UserDialogComponent implements OnInit {
  constructor(
    private readonly fb: FormBuilder,
    public readonly translate: TranslateService,
    public readonly userService: UserService,
    public readonly userList: UserListComponent,
    public readonly loginService: LoginService
  ) {
    this.userForm = this.fb.group({
      name: ["", [Validators.required]],
      orgName: [this.defaultOrg, [Validators.required]],
      // role: ['', [Validators.required]],
      role: [this.defaultRole, [Validators.required]],
      email: ["", [Validators.required, Validators.email]],
      phone: ["", [Validators.required, Validators.pattern("[0-9 ]{11}")]],
    });
    if (loginService.inSiteAdminView()) {
      this.userForm.statusChanges.subscribe((status) => {
        if (this.userForm.value.orgName === "SITE") {
          this.roleTypeList = ["SITE_ADMIN"];
          this.defaultRole = "SITE_ADMIN";
        } else {
          this.roleTypeList = Object.keys(roleTypes4Org);
          this.defaultRole = "ORG_ADMIN";
        }
        // console.log(this.userForm.value.orgName);
      });
    }

    this.orgmap = userList.orgList;
    this.orgNames = Object.keys(userList.orgList);
    this.roleTypeList = Object.keys(roleTypes4Org);
    this.roleType = this.loginService.inSiteAdminView() ? roleTypes : roleTypes4Org;
    this.defaultRole = "ORG_ADMIN";
    this.defaultOrg = this.loginService.inSiteAdminView() ? '':this.loginService.session.orgName;
  }

  @Input() opened = false;
  @Output() save = new EventEmitter();
  @Output() cancel = new EventEmitter();

  userForm: FormGroup;
  orgmap: any;
  orgNames: string[];
  roleTypeList: string[];
  roleType: any;
  defaultRole: string;
  defaultOrg: string;

  readonly vo = {
    serverError: "",
    spinning: false,
  };

  ngOnInit(): void {}

  onCancel() {
    this.close();
  }

  updateOrg(item: string) {
    console.log(item);
    if (item === "SITE" || !item) {
      this.roleTypeList = ["SITE_ADMIN"];
    } else {
      this.roleTypeList = Object.keys(roleTypes4Org);
    }
    return item;
  }

  async onSave() {
    console.log(this.userForm);
    const { value } = this.userForm;
    this.resetModal();
    this.vo.spinning = true;
    await this.userService.addUser(value).then(
      () => {
        this.save.emit("");
        this.close();
        this.vo.spinning = false;
      },
      (error) => {
        this.vo.serverError = error;
        this.vo.spinning = false;
      }
    );
  }

  private close() {
    this.cancel.emit();
  }

  private resetModal() {
    this.vo.serverError = "";
    this.vo.spinning = false;
  }
}
