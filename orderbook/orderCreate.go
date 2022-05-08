package orderbook

import "time"

func Strategy1(asset string, brokerFrom, brokerTo int) OrderContent {
	/*
		/ Strategy 1
		/ > Strictly right after ico announcement
		/ > Ticks go way up
		/ > Short the asset

		/ Type
		/ > High Frequency

		/ Info
		/ > Loss cut at 1.5% * 5(Leverage)
		/ > Draw down at 3% * 5(Leverage)
	*/
	s1 := OrderContent{
		A:   asset,
		N:   Strat1,
		I:   Strat1Id,
		T:   time.Now(),
		ET:  time.Now().Add(time.Second),
		TY:  AssetStrat1,
		SB:  BrokerageInfo[brokerFrom].Name,
		SBC: BrokerageInfo[brokerFrom].Id,
		OB:  BrokerageInfo[brokerTo].Name,
		OBC: BrokerageInfo[brokerTo].Id,
		OD: OrderDetail{
			P:  "limit",
			LV: 5,
			LC: 0.015,
			DD: 0.03,
		},
	}
	return s1
}

func Strategy2(asset string, brokerFrom, brokerTo int) OrderContent {
	/*
		/ Strategy 2
		/ > After some period after ico announcement
		/ > Ticks slowly goes up
		/ > Take long position

		/ Type
		/ > Low Frequency

		/ Info
		/ > Loss cut at 1.5% * 5(Leverage)
		/ > Draw down at 2% * 5(Leverage)
	*/
	s2 := OrderContent{
		A:   asset,
		N:   Strat2,
		I:   Strat2Id,
		T:   time.Now(),
		ET:  time.Now().Add(time.Hour),
		TY:  AssetStrat2,
		SB:  BrokerageInfo[brokerFrom].Name,
		SBC: BrokerageInfo[brokerFrom].Id,
		OB:  BrokerageInfo[brokerTo].Name,
		OBC: BrokerageInfo[brokerTo].Id,
		OD: OrderDetail{
			P:  "limit",
			LV: 5,
			LC: 0.015,
			DD: 0.02,
		},
	}
	return s2
}

func Strategy3(asset string, brokerFrom, brokerTo int) OrderContent {
	/*
		/ Strategy 3
		/ > Right after ico market add
		/ > Ticks goes down
		/ > Take short position

		/ Type
		/ > Low Frequency

		/ Info
		/ > Loss cut at 1.5% * 5(Leverage)
		/ > Draw down at 3% * 5(Leverage)
	*/
	s3 := OrderContent{
		A:   asset,
		N:   Strat3,
		I:   Strat3Id,
		T:   time.Now(),
		ET:  time.Now(),
		TY:  AssetStrat3,
		SB:  BrokerageInfo[brokerFrom].Name,
		SBC: BrokerageInfo[brokerFrom].Id,
		OB:  BrokerageInfo[brokerTo].Name,
		OBC: BrokerageInfo[brokerTo].Id,
		OD: OrderDetail{
			P:  "limit",
			LV: 5,
			LC: 0.01,
			DD: 0.03,
		},
	}
	return s3
}
