const api = {
  // ------------- 设置项 登录
  Login: '/auth/login',// 登录接口
  Logout: '/auth/logout',// 登录接口
  getCaptcha: '/system/imageCaptcha',// 获取验证码
  getSystemInfo: '/system/info',// 获取配置项信息(需要动态路由参数）
 
  // -------------  管理员接口
  getAccount: '/manger/info',// 获取登录管理的信息
  getMangerAuthority: '/manger/authority',// 获取菜单信息
  getMangerList: '/manger/list',// 获取用户列表
  postMangerCreate: '/manger/create',// 创建用户
  getMangerMeta: '/manger/create/meta',// 获取列表
  getMangerEditInfo: '/manger/edit/info',// 获取信息
  postMangerEdit: '/manger/edit',// 修改用户
  postReviewManger: '/manger/review',// 审核
  postRemoveManger: '/manger/remove',// 删除用户

  // ------------- 权限 接口
  getAuthorityList: '/authority/list',// 获取菜单信息
  postaAuthorityCreate: '/authority/create',// 创建菜单
  getAuthorityEditInfo: '/authority/edit/info',// 获取信息
  postAuthorityEdit: '/authority/edit',// 修改
  postAuthorityRemove: '/authority/remove',// 菜单删除
  // ------------- 角色 接口
  getRoleList: '/role/list',// 获取角色列表
  getRoleEditInfo: '/role/edit/info',// 获取信息
  postRoleEdit: '/role/edit',// 修改
  postRoleCreate: '/role/create',// 创建角色
  postRemoveRole: '/role/remove',// 角色删除
  
  // ------------- 设置项 接口
  getSystemSettingInfo: '/system/setting',// 获取配置项信息(需要动态路由参数）
  postSystemSettingSave: '/system/save',// 保存配置项

  // ------------- 分类 接口
  getCateList: '/category/list',// 获取列表
  getCateByModule: '/category/module',// 获取列表
  postCateCreate: '/category/create',// 创建
  getCateEditInfo: '/category/edit/info',// 获取编辑信息
  postCateEdit: '/category/edit',// 编辑
  postCateRemove: '/category/remove',// 删除
  // ------------- 标签 接口
  getTagHots: '/tag/hots',// 获取列表
  getTagList: '/tag/list',// 获取列表
  postTagTop: '/tag/top',// 创建
  postTagRemove: '/tag/remove',// 删除
  // ------------- 媒体 接口
  getMediaList: '/media/list',// 获取列表
  postMediaRemove: '/media/remove',// 删除
  // ------------- 上传文件 接口
  // postuploadFile:  "/media/upload",
  UploadChunk:  "/upload/chunk",
  mergeChunk:  "/upload/mergeChunk",
  // ------------- 幻灯接口
  getCarouselList: '/carousel/list',// 获取列表
  postCarouselCreate: '/carousel/create',// 创建
  postRemoveCarousel: '/carousel/remove',// 删除


  // ------------- 用户 接口
  getUserList: '/user/list',// 获取用户列表
  postUserCreate: '/user/create',// 创建用户
  getUserMeta: '/user/create/meta',// 获取列表
  getUserEditInfo: '/user/edit/info',// 获取信息
  postUserEdit: '/user/edit',// 修改用户
  postReviewUser: '/user/review',// 审核
  postRemoveUser: '/user/remove',// 删除用户


  // ------------- 认证 接口
  getVerifyList: '/verify/list',// 获取列表
  postRemoveVerify: '/verify/remove',// 删除
  postReviewVerify: '/verify/review',// 审核


  
  
  // ------------- 等级 接口
  getGradeList: '/grade/list',// 获取角色列表
  getGradeEditInfo: '/grade/edit/info',// 获取信息
  postGradeEdit: '/grade/edit',// 修改
  postGradeCreate: '/grade/create',// 创建角色
  postRemoveGrade: '/grade/remove',// 角色删除

  // ------------- 会员 接口
  getVipList: '/vip/list',// 获取角色列表
  getVipEditInfo: '/vip/edit/info',// 获取信息
  postVipEdit: '/vip/edit',// 修改
  postVipCreate: '/vip/create',// 创建角色
  postRemoveVip: '/vip/remove',// 角色删除
  


  // ------------- 文章 接口
  getArticleList: '/article/list',// 获取列表
  getArticleMeta: '/article/createMeta',// 获取列表
  postArticleCreate: '/article/create',// 创建
  getArticleEditInfo: '/article/edit/info',// 获取信息
  postArticleEdit: '/article/edit',// 修改
  postRemoveArticle: '/article/remove',// 删除
  postRecoverArticle: '/article/recover',// 移入回收站
  postReductionArticle: '/article/reduction',// 还原
  postReviewArticle: '/article/review',// 审核

  // ------------- 音频 接口
  getAudioList: '/audio/list',// 获取列表
  getAudioListCate: '/audio/list/meta',// 获取列表查询信息
  getAudioMeta: '/audio/create/meta',// 获取列表
  postAudioCreate: '/audio/create',// 创建
  getAudioEditInfo: '/audio/edit/info',// 获取信息
  postAudioEdit: '/audio/edit',// 修改
  postRemoveAudio: '/audio/remove',// 删除
  postRecoverAudio: '/audio/recover',// 移入回收站
  postReductionAudio: '/audio/reduction',// 还原
  postReviewAudio: '/audio/review',// 审核

  // ------------- 视频 接口
  getVideoList: '/video/list',// 获取列表
  getVideoListCate: '/video/list/meta',// 获取列表查询信息
  getVideoMeta: '/video/create/meta',// 获取列表
  postVideoCreate: '/video/create',// 创建
  getVideoEditInfo: '/video/edit/info',// 获取信息
  postVideoEdit: '/video/edit',// 修改
  postRecoverVideo: '/video/recover',// 移入回收站
  postReductionVideo: '/video/reduction',// 还原
  postRemoveVideo: '/video/remove',// 删除
  postReviewVideo: '/video/review',// 审核

  // ------------- 资源 接口
  getResourceList: '/resource/list',// 获取列表
  getResourceListCate: '/resource/list/meta',// 获取列表查询信息
  getResourceMeta: '/resource/create/meta',// 获取列表
  postResourceCreate: '/resource/create',// 创建
  getResourceEditInfo: '/resource/edit/info',// 获取信息
  postResourceEdit: '/resource/edit',// 修改
  postRecoverResource: '/resource/recover',// 移入回收站
  postReductionResource: '/resource/reduction',// 还原
  postRemoveResource: '/resource/remove',// 删除
  postReviewResource: '/resource/review',// 审核

  // ------------- 课程 接口
  getEduList: '/edu/list',// 获取列表
  getEduUserJoinList: '/edu/joinList',// 获取列表
  getEduListCate: '/edu/list/meta',// 获取列表查询信息
  getEduMeta: '/edu/create/meta',// 获取列表
  postEduCreate: '/edu/create',// 创建
  getEduEditInfo: '/edu/edit/info',// 获取信息
  postEduEdit: '/edu/edit',// 修改
  postRecoverEdu: '/edu/recover',// 移入回收站
  postReductionEdu: '/edu/reduction',// 还原
  postRemoveEdu: '/edu/remove',// 删除
  postReviewEdu: '/edu/review',// 审核


  // ------------- 圈子 接口
  getGroupList: '/group/list',// 获取列表
  getGroupListCate: '/group/list/meta',// 获取列表查询信息
  getGroupMeta: '/group/create/meta',// 获取列表
  postGroupCreate: '/group/create',// 创建
  getGroupEditInfo: '/group/edit/info',// 获取信息
  postGroupEdit: '/group/edit',// 修改
  postRecoverGroup: '/group/recover',// 移入回收站
  postReductionGroup: '/group/reduction',// 还原
  postRemoveGroup: '/group/remove',// 删除
  postReviewGroup: '/group/review',// 审核


  // ------------- 问题 接口
  getQuestionList: '/question/list',// 获取列表
  postQuestionCreate: '/question/create',// 创建
  getQuestionEditInfo: '/question/edit/info',// 获取信息
  postQuestionEdit: '/question/edit',// 修改
  postRecoverQuestion: '/question/recover',// 移入回收站
  postReductionQuestion: '/question/reduction',// 还原
  postRemoveQuestion: '/question/remove',// 删除
  postReviewQuestion: '/question/review',// 审核

  // ------------- 话题 接口
  getTopicList: '/topic/list',// 获取列表
  postRecoverTopic: '/topic/recover',// 移入回收站
  postReductionTopic: '/topic/reduction',// 还原
  postRemoveTopic: '/topic/remove',// 删除
  postReviewTopic: '/topic/review',// 审核
  postTopTopic: '/topic/top',// 审核
  // ------------- 评论 接口
  getCommentList: '/comment/list',// 获取列表
  postRecoverComment: '/comment/recover',// 移入回收站
  postReductionComment: '/comment/reduction',// 还原
  postRemoveComment: '/comment/remove',// 删除
  postReviewComment: '/comment/review',// 审核
  // ------------- 回答 接口
  getAnswerList: '/answer/list',// 获取列表
  postRecoverAnswer: '/answer/recover',// 移入回收站
  postReductionAnswer: '/answer/reduction',// 还原
  postRemoveAnswer: '/answer/remove',// 删除
  postReviewAnswer: '/answer/review',// 审核
  // ------------- 举报 接口
  getReportList: '/report/list',// 获取列表
  postRemoveReport: '/report/remove',// 删除
  postReviewReport: '/report/review',// 审核

  // ------------- 订单 接口
  getOrderList: '/order/list',// 获取列表
  getOrderInfo: '/order/info',// 获取列表
  postRemoveOrder: '/order/remove',// 删除
  postReviewOrder: '/order/review',// 审核

  // ------------- 提现 接口
  getCashList: '/cash/list',// 获取列表
  postRemoveCash: '/cash/remove',// 删除
  postReviewCash: '/cash/review',// 审核
  // ------------- 充值 接口
  getRechargeList: '/recharge/list',// 获取列表
  postRemoveRecharge: '/recharge/remove',// 删除
  postReviewRecharge: '/recharge/review',// 审核
  // ------------- 卡密 接口
  getCardList: '/card/list',// 获取列表
  postRemoveCard: '/card/remove',// 删除
  postCreateCard: '/card/create',// 审核


  // ------------- 字典 接口
  getDictType: '/dict/type',// 获取

}
  export default api
  