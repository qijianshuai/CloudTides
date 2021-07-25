export const cloudPlatform = {
  ThinkCloud: 'https://pod.thinkcloud.lenovo.com/api',
  Local: 'http://127.0.0.1:3000/api',
};

export const defaultCloudPlatformURL = cloudPlatform.ThinkCloud;

export const resTypes = {
  Fixed: 'Fixed',
  Dynamic: 'Dynamic',
};

export const roleTypes = {
  SiteAdmin: 'SITE_ADMIN',
  OrgAdmin: 'ORG_ADMIN',
  User: 'USER',
};

export const defaultResType = resTypes.Fixed;
export const defaultRoleType4Site = roleTypes.OrgAdmin;
export const defaultRoleType4Org = roleTypes.User;