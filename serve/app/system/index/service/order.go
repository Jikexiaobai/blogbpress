package service

import (
	"fiber/app/dao"
	"fiber/app/model"
	"fiber/app/system/index/dto"
	"fiber/app/system/index/result"
	"fiber/app/system/index/shared"
	lock_utils "fiber/app/tools/lock"
	"fiber/app/tools/response"
	pay_lib "fiber/library/pay"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/grand"
	"github.com/shopspring/decimal"
)

var Order = new(orderService)

type orderService struct {
}

// SelectList 查询用户订单列表
func (s *orderService) SelectList(req *dto.OrderQuery) (int, []*result.OrderInfo, response.ResponseCode) {
	model := dao.SysOrder.SysOrderDao.Order(dao.SysOrder.Columns.CreateTime, "desc")
	if req.Status != 0 {
		model = model.Where(dao.SysOrder.Columns.Status, req.Status)
	}
	if req.OrderType != 0 {
		model = model.Where(dao.SysOrder.Columns.OrderType, req.OrderType)
	}
	if req.UserId != 0 {
		model = model.Where(dao.SysOrder.Columns.UserId, req.UserId)
	}

	total, err := model.Count()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}

	model = model.Page(req.Page, req.Limit)
	list, err := model.All()
	if err != nil {
		return 0, nil, response.DB_READ_ERROR
	}
	var res []*result.OrderInfo
	for _, i := range list {
		var contentInfo result.OrderInfo
		contentInfo.OrderNum = i.OrderNum
		contentInfo.Money = i.PaymentMoney
		contentInfo.OrderType = i.OrderType
		contentInfo.Status = i.Status
		contentInfo.CreateTime = i.CreateTime
		if i.AuthorId != 0 && i.OrderType != 1 && req.AuthorId == i.AuthorId {
			contentInfo.Money = i.AuthorMoney
			contentInfo.IsIncome = true
		}
		switch i.OrderType {
		case shared.OrderTypeTwo:
			if req.AuthorId == i.AuthorId {
				nickName, err := dao.SysUser.Value(dao.SysUser.Columns.NickName,
					dao.SysUser.Columns.UserId, i.UserId)
				if err != nil {
					return 0, nil, response.DB_READ_ERROR
				}
				contentInfo.Title = "来自" + gconv.String(nickName) + "用户打赏充电"
				contentInfo.DetailId = i.UserId
			} else {
				nickName, err := dao.SysUser.Value(dao.SysUser.Columns.NickName,
					dao.SysUser.Columns.UserId, i.DetailId)
				if err != nil {
					return 0, nil, response.DB_READ_ERROR
				}
				contentInfo.Title = "给" + gconv.String(nickName) + "用户充电打赏"
			}
		case shared.OrderTypeFour:
			if req.AuthorId == i.AuthorId {
				nickName, err := dao.SysUser.Value(dao.SysUser.Columns.NickName,
					dao.SysUser.Columns.UserId, i.UserId)
				if err != nil {
					return 0, nil, response.DB_READ_ERROR
				}
				title, err := dao.SysTopic.Value(dao.SysTopic.Columns.Title, dao.SysTopic.Columns.TopicId, i.DetailId)
				if err != nil {
					return 0, nil, response.DB_READ_ERROR
				}
				contentInfo.Title = gconv.String(nickName) + "付费查看" + gconv.String(title)
				//contentInfo.DetailId = i.UserId
			} else {
				title, err := dao.SysTopic.Value(dao.SysTopic.Columns.Title, dao.SysTopic.Columns.TopicId, i.DetailId)
				if err != nil {
					return 0, nil, response.DB_READ_ERROR
				}
				contentInfo.Title = "我付费查看" + gconv.String(title)
			}
		case shared.OrderTypeThree:
			switch i.DetailModule {
			case shared.Resource:
				if req.AuthorId == i.AuthorId {
					nickName, err := dao.SysUser.Value(dao.SysUser.Columns.NickName,
						dao.SysUser.Columns.UserId, i.UserId)
					if err != nil {
						return 0, nil, response.DB_READ_ERROR
					}
					title, err := dao.SysResource.Value(dao.SysResource.Columns.Title, dao.SysResource.Columns.ResourceId, i.DetailId)
					if err != nil {
						return 0, nil, response.DB_READ_ERROR
					}
					contentInfo.Title = gconv.String(nickName) + "付费购买" + gconv.String(title)
					//contentInfo.DetailId = i.UserId
				} else {
					title, err := dao.SysResource.Value(dao.SysResource.Columns.Title, dao.SysResource.Columns.ResourceId, i.DetailId)
					if err != nil {
						return 0, nil, response.DB_READ_ERROR
					}
					contentInfo.Title = "我付费购买" + gconv.String(title)
				}
			case shared.Audio:
				if req.AuthorId == i.AuthorId {
					nickName, err := dao.SysUser.Value(dao.SysUser.Columns.NickName,
						dao.SysUser.Columns.UserId, i.UserId)
					if err != nil {
						return 0, nil, response.DB_READ_ERROR
					}
					title, err := dao.SysAudio.Value(dao.SysAudio.Columns.Title, dao.SysAudio.Columns.AudioId, i.DetailId)
					if err != nil {
						return 0, nil, response.DB_READ_ERROR
					}
					contentInfo.Title = gconv.String(nickName) + "付费购买" + gconv.String(title)
					//contentInfo.DetailId = i.UserId
				} else {
					title, err := dao.SysAudio.Value(dao.SysAudio.Columns.Title, dao.SysAudio.Columns.AudioId, i.DetailId)
					if err != nil {
						return 0, nil, response.DB_READ_ERROR
					}
					contentInfo.Title = "我付费购买" + gconv.String(title)
				}
			case shared.Video:
				if req.AuthorId == i.AuthorId {
					nickName, err := dao.SysUser.Value(dao.SysUser.Columns.NickName,
						dao.SysUser.Columns.UserId, i.UserId)
					if err != nil {
						return 0, nil, response.DB_READ_ERROR
					}
					title, err := dao.SysVideo.Value(dao.SysVideo.Columns.Title, dao.SysVideo.Columns.VideoId, i.DetailId)
					if err != nil {
						return 0, nil, response.DB_READ_ERROR
					}
					contentInfo.Title = gconv.String(nickName) + "付费购买" + gconv.String(title)
					//contentInfo.DetailId = i.UserId
				} else {
					title, err := dao.SysVideo.Value(dao.SysVideo.Columns.Title, dao.SysVideo.Columns.VideoId, i.DetailId)
					if err != nil {
						return 0, nil, response.DB_READ_ERROR
					}
					contentInfo.Title = "我付费购买" + gconv.String(title)
				}
			}
		case shared.OrderTypeFive:
			if req.AuthorId == i.AuthorId {
				nickName, err := dao.SysUser.Value(dao.SysUser.Columns.NickName,
					dao.SysUser.Columns.UserId, i.UserId)
				if err != nil {
					return 0, nil, response.DB_READ_ERROR
				}
				title, err := dao.SysGroup.Value(dao.SysGroup.Columns.Title, dao.SysGroup.Columns.GroupId, i.DetailId)
				if err != nil {
					return 0, nil, response.DB_READ_ERROR
				}
				contentInfo.Title = gconv.String(nickName) + "付费加入" + gconv.String(title)
				//contentInfo.DetailId = i.UserId
			} else {
				title, err := dao.SysGroup.Value(dao.SysGroup.Columns.Title, dao.SysGroup.Columns.GroupId, i.DetailId)
				if err != nil {
					return 0, nil, response.DB_READ_ERROR
				}
				contentInfo.Title = "我付费加入" + gconv.String(title)
			}
		case shared.OrderTypeSix:
			if req.AuthorId == i.AuthorId {
				nickName, err := dao.SysUser.Value(dao.SysUser.Columns.NickName,
					dao.SysUser.Columns.UserId, i.UserId)
				if err != nil {
					return 0, nil, response.DB_READ_ERROR
				}
				title, err := dao.SysEdu.Value(dao.SysEdu.Columns.Title, dao.SysEdu.Columns.EduId, i.DetailId)
				if err != nil {
					return 0, nil, response.DB_READ_ERROR
				}
				contentInfo.Title = gconv.String(nickName) + "付费加入" + gconv.String(title)
				//contentInfo.DetailId = i.UserId
			} else {
				title, err := dao.SysEdu.Value(dao.SysEdu.Columns.Title, dao.SysEdu.Columns.EduId, i.DetailId)
				if err != nil {
					return 0, nil, response.DB_READ_ERROR
				}
				contentInfo.Title = "我付费加入" + gconv.String(title)
			}
		case shared.OrderTypeSeven:
			if req.AuthorId == i.AuthorId {
				nickName, err := dao.SysUser.Value(dao.SysUser.Columns.NickName,
					dao.SysUser.Columns.UserId, i.UserId)
				if err != nil {
					return 0, nil, response.DB_READ_ERROR
				}
				questionId, err := dao.SysAnswer.Value(dao.SysAnswer.Columns.TopicId, dao.SysAnswer.Columns.AnswerId, i.DetailId)
				if err != nil {
					return 0, nil, response.DB_READ_ERROR
				}
				title, err := dao.SysQuestion.Value(dao.SysQuestion.Columns.Title, dao.SysQuestion.Columns.QuestionId, questionId)
				if err != nil {
					return 0, nil, response.DB_READ_ERROR
				}
				contentInfo.Title = gconv.String(nickName) + "付费查看" + gconv.String(title)
				//contentInfo.DetailId = i.UserId
			} else {
				questionId, err := dao.SysAnswer.Value(dao.SysAnswer.Columns.TopicId, dao.SysAnswer.Columns.AnswerId, i.DetailId)
				if err != nil {
					return 0, nil, response.DB_READ_ERROR
				}
				title, err := dao.SysQuestion.Value(dao.SysQuestion.Columns.Title, dao.SysQuestion.Columns.QuestionId, questionId)
				if err != nil {
					return 0, nil, response.DB_READ_ERROR
				}
				contentInfo.Title = "我付费查看" + gconv.String(title)
			}
		case shared.OrderTypeEight:
			title, err := dao.SysRole.Value(dao.SysRole.Columns.Title, dao.SysRole.Columns.RoleId, i.DetailId)
			if err != nil {
				return 0, nil, response.DB_READ_ERROR
			}
			contentInfo.Title = "我付费开通了" + gconv.String(title)
		case shared.OrderTypeNine:
			contentInfo.Title = "支付认证服务费"
		}
		res = append(res, &contentInfo)
	}

	return total, res, response.SUCCESS
}

// Create 创建订单
func (s *orderService) Create(req *dto.OrderCreate) (string, response.ResponseCode) {
	// 加入锁限制
	_, err := lock_utils.SetCount(shared.OrderCreateCount+gconv.String(req.UserId),
		shared.OrderCreateLock+gconv.String(req.UserId), 60, 5)
	if err != nil {
		return "", response.CACHE_SAVE_ERROR
	}

	var order model.SysOrder
	err = gconv.Struct(req, &order)
	if err != nil {
		return "", response.INVALID
	}
	// 生成流水号
	OrderNum := s.buildNumber(req.OrderType, req.PayMethod, req.UserId)
	order.OrderNum = OrderNum
	order.Status = 1
	order.CreateTime = gtime.Now()
	order.UpdateTime = gtime.Now()
	order.OrderMode = 1
	order.UserId = req.UserId

	// 获取折扣
	discount, err := Vip.SelectVipDiscount(req.UserId)
	if err != nil {
		return "", response.DB_READ_ERROR
	}

	// 获取分成比例
	paySetting, err := Config.FindValue(shared.PaySetting)
	if err != nil {
		return "", response.DB_READ_ERROR
	}
	payJson := gjson.New(paySetting)
	servicePercent := gconv.Float64(payJson.Get("servicePercent"))

	switch req.OrderType {
	case shared.OrderTypeTwo:
		order.AuthorId = req.DetailId
		orderMoney := decimal.NewFromFloat(gconv.Float64(req.OrderMoney))
		if servicePercent != 0 {
			order.ServiceMoney = gconv.Float64(orderMoney.Mul(decimal.NewFromFloat(servicePercent)))
			order.PaymentMoney = gconv.Float64(orderMoney)
			order.AuthorMoney = gconv.Float64(orderMoney.Sub(decimal.NewFromFloat(order.ServiceMoney)))
		} else {
			order.PaymentMoney = gconv.Float64(orderMoney)
			order.AuthorMoney = order.PaymentMoney
		}
	case shared.OrderTypeFour:
		info, err := dao.SysTopic.
			Fields(dao.SysTopic.Columns.Price, dao.SysTopic.Columns.UserId).
			Where(dao.SysTopic.Columns.TopicId, req.DetailId).
			One(dao.SysTopic.Columns.Price)
		if err != nil {
			return "", response.DB_READ_ERROR
		}

		orderMoney := decimal.NewFromFloat(gconv.Float64(info.Price))
		order.OrderMoney = gconv.Float64(orderMoney)
		order.AuthorId = info.UserId

		if discount != 0 {
			districtMoney := orderMoney.Mul(decimal.NewFromFloat(gconv.Float64(discount)))
			paymentMoney := orderMoney.Sub(districtMoney)
			order.PaymentMoney = gconv.Float64(paymentMoney)
			order.DistrictMoney = gconv.Float64(districtMoney)
			if servicePercent != 0 {
				order.ServiceMoney = gconv.Float64(paymentMoney.Mul(decimal.NewFromFloat(servicePercent)))
				order.AuthorMoney = gconv.Float64(paymentMoney.Sub(decimal.NewFromFloat(order.ServiceMoney)))
			} else {
				order.AuthorMoney = order.PaymentMoney
			}

		} else {
			if servicePercent != 0 {
				order.ServiceMoney = gconv.Float64(orderMoney.Mul(decimal.NewFromFloat(servicePercent)))
				order.PaymentMoney = gconv.Float64(orderMoney)
				order.AuthorMoney = gconv.Float64(orderMoney.Sub(decimal.NewFromFloat(order.ServiceMoney)))
			} else {
				order.PaymentMoney = gconv.Float64(orderMoney)
				order.AuthorMoney = order.PaymentMoney
			}

		}
	case shared.OrderTypeThree:
		switch req.DetailModule {
		case shared.Resource:
			info, err := dao.SysResource.Fields(dao.SysResource.Columns.Price, dao.SysResource.Columns.UserId).Where(dao.SysResource.Columns.ResourceId, req.DetailId).One()
			if err != nil {
				return "", response.DB_READ_ERROR
			}
			orderMoney := decimal.NewFromFloat(gconv.Float64(info.Price))
			order.OrderMoney = gconv.Float64(orderMoney)
			order.AuthorId = info.UserId

			if discount != 0 {
				districtMoney := orderMoney.Mul(decimal.NewFromFloat(gconv.Float64(discount)))
				paymentMoney := orderMoney.Sub(districtMoney)
				order.PaymentMoney = gconv.Float64(paymentMoney)
				order.DistrictMoney = gconv.Float64(districtMoney)
				if servicePercent != 0 {
					order.ServiceMoney = gconv.Float64(paymentMoney.Mul(decimal.NewFromFloat(servicePercent)))
					order.AuthorMoney = gconv.Float64(paymentMoney.Sub(decimal.NewFromFloat(order.ServiceMoney)))
				} else {
					order.AuthorMoney = order.PaymentMoney
				}

			} else {
				if servicePercent != 0 {
					order.ServiceMoney = gconv.Float64(orderMoney.Mul(decimal.NewFromFloat(servicePercent)))
					order.PaymentMoney = gconv.Float64(orderMoney)
					order.AuthorMoney = gconv.Float64(orderMoney.Sub(decimal.NewFromFloat(order.ServiceMoney)))
				} else {
					order.PaymentMoney = gconv.Float64(orderMoney)
					order.AuthorMoney = order.PaymentMoney
				}

			}
		case shared.Audio:
			info, err := dao.SysAudio.Fields(dao.SysAudio.Columns.Price, dao.SysAudio.Columns.UserId).Where(dao.SysAudio.Columns.AudioId, req.DetailId).One()
			if err != nil {
				return "", response.DB_READ_ERROR
			}

			orderMoney := decimal.NewFromFloat(gconv.Float64(info.Price))
			order.OrderMoney = gconv.Float64(orderMoney)
			order.AuthorId = info.UserId
			if discount != 0 {
				districtMoney := orderMoney.Mul(decimal.NewFromFloat(gconv.Float64(discount)))
				paymentMoney := orderMoney.Sub(districtMoney)
				order.PaymentMoney = gconv.Float64(paymentMoney)
				order.DistrictMoney = gconv.Float64(districtMoney)
				if servicePercent != 0 {
					order.ServiceMoney = gconv.Float64(paymentMoney.Mul(decimal.NewFromFloat(servicePercent)))
					order.AuthorMoney = gconv.Float64(paymentMoney.Sub(decimal.NewFromFloat(order.ServiceMoney)))
				} else {
					order.AuthorMoney = order.PaymentMoney
				}

			} else {
				if servicePercent != 0 {
					order.ServiceMoney = gconv.Float64(orderMoney.Mul(decimal.NewFromFloat(servicePercent)))
					order.PaymentMoney = gconv.Float64(orderMoney)
					order.AuthorMoney = gconv.Float64(orderMoney.Sub(decimal.NewFromFloat(order.ServiceMoney)))
				} else {
					order.PaymentMoney = gconv.Float64(orderMoney)
					order.AuthorMoney = order.PaymentMoney
				}

			}
		case shared.Video:
			info, err := dao.SysVideo.
				Fields(dao.SysVideo.Columns.Price, dao.SysVideo.Columns.UserId).
				Where(dao.SysVideo.Columns.VideoId, req.DetailId).One()
			if err != nil {
				return "", response.DB_READ_ERROR
			}

			orderMoney := decimal.NewFromFloat(gconv.Float64(info.Price))
			order.OrderMoney = gconv.Float64(orderMoney)
			order.AuthorId = info.UserId

			if discount != 0 {
				districtMoney := orderMoney.Mul(decimal.NewFromFloat(gconv.Float64(discount)))
				paymentMoney := orderMoney.Sub(districtMoney)
				order.PaymentMoney = gconv.Float64(paymentMoney)
				order.DistrictMoney = gconv.Float64(districtMoney)
				if servicePercent != 0 {
					order.ServiceMoney = gconv.Float64(paymentMoney.Mul(decimal.NewFromFloat(servicePercent)))
					order.AuthorMoney = gconv.Float64(paymentMoney.Sub(decimal.NewFromFloat(order.ServiceMoney)))
				} else {
					order.AuthorMoney = order.PaymentMoney
				}

			} else {
				if servicePercent != 0 {
					order.ServiceMoney = gconv.Float64(orderMoney.Mul(decimal.NewFromFloat(servicePercent)))
					order.PaymentMoney = gconv.Float64(orderMoney)
					order.AuthorMoney = gconv.Float64(orderMoney.Sub(decimal.NewFromFloat(order.ServiceMoney)))
				} else {
					order.PaymentMoney = gconv.Float64(orderMoney)
					order.AuthorMoney = order.PaymentMoney
				}

			}
		}
	case shared.OrderTypeFive:

		info, err := dao.SysGroup.
			Fields(dao.SysGroup.Columns.Price, dao.SysGroup.Columns.UserId).
			Where(dao.SysGroup.Columns.GroupId, req.DetailId).One()
		if err != nil {
			return "", response.DB_READ_ERROR
		}

		orderMoney := decimal.NewFromFloat(gconv.Float64(info.Price))
		order.OrderMoney = gconv.Float64(orderMoney)
		order.AuthorId = info.UserId
		if discount != 0 {
			districtMoney := orderMoney.Mul(decimal.NewFromFloat(gconv.Float64(discount)))
			paymentMoney := orderMoney.Sub(districtMoney)
			order.PaymentMoney = gconv.Float64(paymentMoney)
			order.DistrictMoney = gconv.Float64(districtMoney)
			if servicePercent != 0 {
				order.ServiceMoney = gconv.Float64(paymentMoney.Mul(decimal.NewFromFloat(servicePercent)))
				order.AuthorMoney = gconv.Float64(paymentMoney.Sub(decimal.NewFromFloat(order.ServiceMoney)))
			} else {
				order.AuthorMoney = order.PaymentMoney
			}

		} else {
			if servicePercent != 0 {
				order.ServiceMoney = gconv.Float64(orderMoney.Mul(decimal.NewFromFloat(servicePercent)))
				order.PaymentMoney = gconv.Float64(orderMoney)
				order.AuthorMoney = gconv.Float64(orderMoney.Sub(decimal.NewFromFloat(order.ServiceMoney)))
			} else {
				order.PaymentMoney = gconv.Float64(orderMoney)
				order.AuthorMoney = order.PaymentMoney
			}

		}
	case shared.OrderTypeSix:
		info, err := dao.SysEdu.
			Fields(dao.SysEdu.Columns.Price, dao.SysEdu.Columns.UserId).
			Where(dao.SysEdu.Columns.EduId, req.DetailId).One()
		if err != nil {
			return "", response.DB_READ_ERROR
		}

		orderMoney := decimal.NewFromFloat(gconv.Float64(info.Price))
		order.OrderMoney = gconv.Float64(orderMoney)
		order.AuthorId = info.UserId
		order.AuthorMoney = gconv.Float64(orderMoney)
		if discount != 0 {
			districtMoney := orderMoney.Mul(decimal.NewFromFloat(gconv.Float64(discount)))
			paymentMoney := orderMoney.Sub(districtMoney)
			order.PaymentMoney = gconv.Float64(paymentMoney)
			order.DistrictMoney = gconv.Float64(districtMoney)
			if servicePercent != 0 {
				order.ServiceMoney = gconv.Float64(paymentMoney.Mul(decimal.NewFromFloat(servicePercent)))
				order.AuthorMoney = gconv.Float64(paymentMoney.Sub(decimal.NewFromFloat(order.ServiceMoney)))
			} else {
				order.AuthorMoney = order.PaymentMoney
			}

		} else {
			if servicePercent != 0 {
				order.ServiceMoney = gconv.Float64(orderMoney.Mul(decimal.NewFromFloat(servicePercent)))
				order.PaymentMoney = gconv.Float64(orderMoney)
				order.AuthorMoney = gconv.Float64(orderMoney.Sub(decimal.NewFromFloat(order.ServiceMoney)))
			} else {
				order.PaymentMoney = gconv.Float64(orderMoney)
				order.AuthorMoney = order.PaymentMoney
			}

		}
	case shared.OrderTypeSeven:
		info, err := dao.SysAnswer.
			Fields(dao.SysAnswer.Columns.Price, dao.SysAnswer.Columns.UserId).
			Where(dao.SysAnswer.Columns.AnswerId, req.DetailId).One()
		if err != nil {
			return "", response.DB_READ_ERROR
		}

		orderMoney := decimal.NewFromFloat(gconv.Float64(info.Price))
		order.OrderMoney = gconv.Float64(orderMoney)
		order.AuthorId = info.UserId
		order.AuthorMoney = gconv.Float64(orderMoney)
		if discount != 0 {
			districtMoney := orderMoney.Mul(decimal.NewFromFloat(gconv.Float64(discount)))
			paymentMoney := orderMoney.Sub(districtMoney)
			order.PaymentMoney = gconv.Float64(paymentMoney)
			order.DistrictMoney = gconv.Float64(districtMoney)
			if servicePercent != 0 {
				order.ServiceMoney = gconv.Float64(paymentMoney.Mul(decimal.NewFromFloat(servicePercent)))
				order.AuthorMoney = gconv.Float64(paymentMoney.Sub(decimal.NewFromFloat(order.ServiceMoney)))
			} else {
				order.AuthorMoney = order.PaymentMoney
			}

		} else {
			if servicePercent != 0 {
				order.ServiceMoney = gconv.Float64(orderMoney.Mul(decimal.NewFromFloat(servicePercent)))
				order.PaymentMoney = gconv.Float64(orderMoney)
				order.AuthorMoney = gconv.Float64(orderMoney.Sub(decimal.NewFromFloat(order.ServiceMoney)))
			} else {
				order.PaymentMoney = gconv.Float64(orderMoney)
				order.AuthorMoney = order.PaymentMoney
			}

		}
	case shared.OrderTypeEight:
		price, err := dao.SysVip.Value(dao.SysVip.Columns.Price, dao.SysVip.Columns.VipId, req.DetailId)
		if err != nil {
			return "", response.DB_READ_ERROR
		}
		orderMoney := decimal.NewFromFloat(gconv.Float64(price))
		order.OrderMoney = gconv.Float64(orderMoney)
		order.PaymentMoney = gconv.Float64(orderMoney)
	case shared.OrderTypeNine:
		verifySetting, err := Config.FindValue(shared.UserSetting)
		if err != nil {
			return "", response.DB_READ_ERROR
		}
		j := gjson.New(verifySetting)
		price := gconv.Float64(j.Get("verifyPrice"))
		orderMoney := decimal.NewFromFloat(price)
		order.OrderMoney = gconv.Float64(orderMoney)
		order.PaymentMoney = gconv.Float64(orderMoney)
	}

	_, err = dao.SysOrder.Save(order)
	if err != nil {
		return "", response.ADD_FAILED
	}

	return OrderNum, response.SUCCESS
}

// Pay 支付订单
func (s *orderService) Pay(userId int64, orderNum string) (qrInfo *result.OrderPayInfo, code response.ResponseCode) {

	order, err := dao.SysOrder.
		Where(dao.SysOrder.Columns.UserId, userId).
		Where(dao.SysOrder.Columns.OrderNum, orderNum).One()
	if err != nil || order == nil {
		return nil, response.NOT_FOUND
	}

	if order.PayMethod == shared.OrderBalance {
		tx, err := g.DB().Begin()
		if err != nil {
			return nil, response.FAILD
		}
		defer func() {
			if code != response.SUCCESS {
				tx.Rollback()
			} else {
				tx.Commit()
			}
		}()
		// 自己的余额
		userBalance, err := Account.SelectBalance(order.UserId)
		if err != nil {
			return nil, response.DB_READ_ERROR
		}
		// 比较余额
		cmp := decimal.NewFromFloat(order.PaymentMoney).
			Cmp(decimal.NewFromFloat(userBalance))
		if cmp == 1 {
			return nil, response.INVALID
		}

		//修改余额
		balance := decimal.NewFromFloat(userBalance).Sub(decimal.NewFromFloat(order.PaymentMoney))
		err = Account.EditBalance(tx, order.UserId, gconv.Float64(balance))
		if err != nil {
			return nil, response.UPDATE_FAILED
		}

		// 修改订单信息
		_, err = tx.Update(dao.SysOrder.Table, g.Map{
			dao.SysOrder.Columns.Status:  shared.StatusReviewed,
			dao.SysOrder.Columns.PayTime: gtime.Now(),
		}, dao.SysOrder.Columns.OrderId, order.OrderId)
		if err != nil {
			return nil, response.UPDATE_FAILED
		}

		// 设置作者的余额并且通知作者
		if order.OrderType != shared.OrderTypeEight && order.OrderType != shared.OrderTypeNine {
			// 设置作者余额
			authorBalance, err := Account.SelectBalance(order.AuthorId)
			if err != nil {
				return nil, response.DB_READ_ERROR
			}

			tmpBalance := decimal.NewFromFloat(authorBalance).
				Add(decimal.NewFromFloat(order.AuthorMoney))
			err = Account.EditBalance(tx, order.AuthorId, gconv.Float64(tmpBalance))
			if err != nil {
				return nil, response.UPDATE_FAILED
			}

			//	通知作者
			err = Account.EditBalance(tx, order.AuthorId, gconv.Float64(tmpBalance))
			if err != nil {
				return nil, response.UPDATE_FAILED
			}
			baseSetting, err := Config.FindValue(shared.BaseSetting)
			if err != nil {
				return nil, response.DB_READ_ERROR
			}
			// 获取文件存储设置
			j := gjson.New(baseSetting)
			currencySymbol := gconv.String(j.Get("currencySymbol"))
			var notice model.SysNotice
			switch order.OrderType {
			case shared.OrderTypeTwo:
				nickName, err := dao.SysUser.Value(dao.SysUser.Columns.NickName, dao.SysUser.Columns.UserId, order.UserId)
				if err != nil {
					return nil, response.UPDATE_FAILED
				}

				notice.Content = gconv.String(nickName) + "给您打赏了" + gconv.String(order.PaymentMoney) + currencySymbol
				notice.SystemType = shared.NoticeUserTips
				notice.DetailId = order.OrderId
			case shared.OrderTypeThree:
				nickName, err := dao.SysUser.Value(dao.SysUser.Columns.NickName, dao.SysUser.Columns.UserId, order.UserId)
				if err != nil {
					return nil, response.UPDATE_FAILED
				}
				var title string
				switch order.DetailModule {
				case shared.Resource:
					tmpTitle, err := dao.SysResource.
						Value(dao.SysResource.Columns.Title, dao.SysResource.Columns.ResourceId, order.DetailId)
					if err != nil {
						return nil, response.UPDATE_FAILED
					}
					title = gconv.String(tmpTitle)
				case shared.Video:
					tmpTitle, err := dao.SysVideo.
						Value(dao.SysVideo.Columns.Title, dao.SysVideo.Columns.VideoId, order.DetailId)
					if err != nil {
						return nil, response.UPDATE_FAILED
					}
					title = gconv.String(tmpTitle)
				case shared.Audio:
					tmpTitle, err := dao.SysAudio.
						Value(dao.SysAudio.Columns.Title, dao.SysAudio.Columns.AudioId, order.DetailId)
					if err != nil {
						return nil, response.UPDATE_FAILED
					}
					title = gconv.String(tmpTitle)
				}
				notice.Content = gconv.String(nickName) + "花费" + gconv.String(order.PaymentMoney) + currencySymbol + "购买了《" + title + "》里的资源内容"
				notice.SystemType = shared.NoticeUserBuyContent
				notice.DetailId = order.OrderId
			}
			notice.Type = shared.NoticeSystem
			notice.Receiver = order.AuthorId
			notice.CreateTime = gtime.Now()
			notice.Status = shared.NoticeStatusReview

			_, err = tx.Insert(dao.SysNotice.Table, notice)
			if err != nil {
				return nil, response.DB_SAVE_ERROR
			}
		}

		// 设置用户开通会员
		if order.OrderType == shared.OrderTypeNine {
			err = Vip.OpenVip(tx, order.DetailId, order.UserId)
			if err != nil {
				return nil, response.UPDATE_FAILED
			}
		}
		qrInfo = &result.OrderPayInfo{
			PayMethod: order.PayMethod,
			IsPay:     true,
			OrderNum:  orderNum,
		}
	} else {

		payOptions, err := Config.FindValue(shared.PaySetting)
		if err != nil {
			return nil, response.DB_READ_ERROR
		}
		// 获取文件存储设置
		j := gjson.New(payOptions)
		alyPay := gconv.String(j.Get("alyPay"))
		weChatPay := gconv.String(j.Get("weChatPay"))
		orderNotice := gconv.String(j.Get("orderNotice"))

		switch order.PayMethod {
		case shared.OrderAliPay:
			switch alyPay {
			case "close":
				return nil, response.PAY_ERROR
			case "aly":
				alyOptions, err := Config.FindValue("AlyPayOptions")
				if err != nil {
					return nil, response.PAY_ERROR
				}
				// 获取文件存储设置
				j := gjson.New(alyOptions)
				appId := gconv.String(j.Get("appId"))
				privateKey := gconv.String(j.Get("privateKey"))

				alyCertPublicKey := Media.SelectMediaPath(gconv.String(j.Get("alyCertPublicKey")))
				alyRootCert := Media.SelectMediaPath(gconv.String(j.Get("alyRootCert")))
				appPublicKey := Media.SelectMediaPath(gconv.String(j.Get("appPublicKey")))
				payEngine := &pay_lib.AliPayBody{
					AppId:            appId,
					PrivateKey:       privateKey,
					NotifyUrl:        orderNotice,
					Subject:          s.restOrderTitle(order.OrderType),
					OutTradeNo:       order.OrderNum,
					TotalAmount:      order.PaymentMoney,
					AlyCertPublicKey: "." + alyCertPublicKey,
					AlyRootCert:      "." + alyRootCert,
					AppCertPublicKey: "." + appPublicKey,
				}
				pageUrl, err := payEngine.WebPay()
				if err != nil {
					return nil, response.PAY_ERROR
				}

				qrInfo = &result.OrderPayInfo{
					PayMethod: order.PayMethod,
					QrCode:    pageUrl,
					OrderNum:  orderNum,
				}
			}
		case shared.OrderWechatPay:
			switch weChatPay {
			case "close":
				return nil, response.PAY_ERROR
			case "weChat":
				return nil, response.PAY_ERROR
			}
		}
	}
	return qrInfo, response.SUCCESS
}

// CheckPay 检查订单是否支付
func (s *orderService) CheckPay(userId int64, orderNum string) bool {
	res, err := dao.SysOrder.
		Where(dao.SysOrder.Columns.UserId, userId).
		Value(dao.SysOrder.Columns.Status,
			dao.SysOrder.Columns.OrderNum, orderNum)
	if err != nil || gconv.Int(res) == 2 {
		return true
	}
	return false
}

// CheckIsPay 查询订单内容订单状态
func (s *orderService) CheckIsPay(userId, productId int64, module string, orderType int) bool {
	model := dao.SysOrder.SysOrderDao.
		Where(dao.SysOrder.Columns.UserId, userId)
	if productId != 0 {
		model = model.Where(dao.SysOrder.Columns.DetailId, productId)
	}
	if module != "" {
		model = model.Where(dao.SysOrder.Columns.DetailModule, module)
	}
	if orderType != 0 {
		model = model.Where(dao.SysOrder.Columns.OrderType, orderType)
	}
	list, err := model.All()
	if err != nil || len(list) < 1 {
		return false
	}
	for _, i := range list {
		if i.Status == 2 {
			return true
		}
	}
	return true
}

// UpdateStatus 更新订单状态
func (s *orderService) UpdateStatus(orderNum string) (code response.ResponseCode) {

	order, err := dao.SysOrder.Where(dao.SysOrder.Columns.OrderNum, orderNum).One()
	if err != nil || order == nil {
		return response.NOT_FOUND
	}
	tx, err := g.DB().Begin()
	if err != nil {
		return response.DB_TX_ERROR
	}

	defer func() {
		if code != response.SUCCESS {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// 设置作者的余额
	if order.OrderType != shared.OrderTypeEight && order.OrderType != shared.OrderTypeNine {
		// 设置作者余额
		authorBalance, err := Account.SelectBalance(order.AuthorId)
		if err != nil {
			return response.DB_READ_ERROR
		}

		tmpBalance := decimal.NewFromFloat(authorBalance).
			Add(decimal.NewFromFloat(order.AuthorMoney))
		err = Account.EditBalance(tx, order.AuthorId, gconv.Float64(tmpBalance))
		if err != nil {
			return response.UPDATE_FAILED
		}

		baseSetting, err := Config.FindValue(shared.BaseSetting)
		if err != nil {
			return response.UPDATE_FAILED
		}
		// 获取文件存储设置
		j := gjson.New(baseSetting)
		currencySymbol := gconv.String(j.Get("currencySymbol"))
		var notice model.SysNotice
		switch order.OrderType {
		case shared.OrderTypeTwo:
			nickName, err := dao.SysUser.Value(dao.SysUser.Columns.NickName, dao.SysUser.Columns.UserId, order.UserId)
			if err != nil {
				return response.UPDATE_FAILED
			}

			notice.Content = gconv.String(nickName) + "给您打赏了" + gconv.String(order.PaymentMoney) + currencySymbol
			notice.SystemType = shared.NoticeUserTips
		case shared.OrderTypeThree:
			nickName, err := dao.SysUser.Value(dao.SysUser.Columns.NickName, dao.SysUser.Columns.UserId, order.UserId)
			if err != nil {
				return response.UPDATE_FAILED
			}
			var title string
			switch order.DetailModule {
			case shared.Resource:
				tmpTitle, err := dao.SysResource.
					Value(dao.SysResource.Columns.Title, dao.SysResource.Columns.ResourceId, order.DetailId)
				if err != nil {
					return response.UPDATE_FAILED
				}
				title = gconv.String(tmpTitle)
			case shared.Video:
				tmpTitle, err := dao.SysVideo.
					Value(dao.SysVideo.Columns.Title, dao.SysVideo.Columns.VideoId, order.DetailId)
				if err != nil {
					return response.UPDATE_FAILED
				}
				title = gconv.String(tmpTitle)
			case shared.Audio:
				tmpTitle, err := dao.SysAudio.
					Value(dao.SysAudio.Columns.Title, dao.SysAudio.Columns.AudioId, order.DetailId)
				if err != nil {
					return response.UPDATE_FAILED
				}
				title = gconv.String(tmpTitle)
			}
			notice.Content = gconv.String(nickName) + "花费" + gconv.String(order.PaymentMoney) + currencySymbol + "购买了《" + title + "》里的资源内容"
			notice.SystemType = shared.NoticeUserBuyContent
		}
		notice.Type = shared.NoticeSystem
		notice.Receiver = order.AuthorId
		notice.CreateTime = gtime.Now()
		notice.Status = shared.NoticeStatusReview

		_, err = tx.Insert(dao.SysNotice.Table, notice)
		if err != nil {
			return response.DB_SAVE_ERROR
		}

	}

	// 设置用户开通会员
	if order.OrderType == shared.OrderTypeNine {
		err = Vip.OpenVip(tx, order.DetailId, order.UserId)
		if err != nil {
			return response.UPDATE_FAILED
		}
	}

	_, err = dao.SysOrder.Data(g.Map{
		dao.SysOrder.Columns.Status:  1,
		dao.SysOrder.Columns.PayTime: gtime.Now(),
	}).Where(dao.SysOrder.Columns.OrderNum, orderNum).Update()
	if err != nil {
		return response.UPDATE_FAILED
	}

	return response.SUCCESS
}

func (s *orderService) restOrderType(orderType int) string {
	msg := map[int]string{
		shared.OrderTypeThree: "Z",
		shared.OrderTypeTwo:   "C",
		shared.OrderTypeFour:  "V",
		shared.OrderTypeFive:  "J",
		shared.OrderTypeSix:   "J",
		shared.OrderTypeSeven: "VA",
		shared.OrderTypeEight: "V",
		shared.OrderTypeNine:  "V",
	}
	return msg[orderType]
}

func (s *orderService) restOrderTitle(orderType int) string {
	msg := map[int]string{
		shared.OrderTypeThree: "购买资源",
		shared.OrderTypeTwo:   "用户打赏",
		shared.OrderTypeFour:  "购买隐藏内容",
		shared.OrderTypeFive:  "加入付费圈子",
		shared.OrderTypeSix:   "加入付费课程",
		shared.OrderTypeSeven: "查看付费答案",
		shared.OrderTypeEight: "开通会员",
		shared.OrderTypeNine:  "认证付费",
	}
	return msg[orderType]
}

//BuildNumber 业务类型 + 毫秒时间戳 + 6位随机数 + 用户id
func (s *orderService) buildNumber(orderType, paymentMethod int, userId int64) string {
	time := gconv.String(gtime.TimestampMilli())
	rand := grand.Digits(6)
	return s.restOrderType(orderType) + gconv.String(orderType) + time + rand + gconv.String(paymentMethod) + gconv.String(userId)
}
