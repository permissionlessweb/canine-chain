import { Client, registry, MissingWalletError } from 'jackal-dao-canine-client-ts'

import { LPool } from "jackal-dao-canine-client-ts/jackaldao.canine.lp/types"
import { LProviderRecord } from "jackal-dao-canine-client-ts/jackaldao.canine.lp/types"
import { Params } from "jackal-dao-canine-client-ts/jackaldao.canine.lp/types"


export { LPool, LProviderRecord, Params };

function initClient(vuexGetters) {
	return new Client(vuexGetters['common/env/getEnv'], vuexGetters['common/wallet/signer'])
}

function mergeResults(value, next_values) {
	for (let prop of Object.keys(next_values)) {
		if (Array.isArray(next_values[prop])) {
			value[prop]=[...value[prop], ...next_values[prop]]
		}else{
			value[prop]=next_values[prop]
		}
	}
	return value
}

type Field = {
	name: string;
	type: unknown;
}
function getStructure(template) {
	let structure: {fields: Field[]} = { fields: [] }
	for (const [key, value] of Object.entries(template)) {
		let field = { name: key, type: typeof value }
		structure.fields.push(field)
	}
	return structure
}
const getDefaultState = () => {
	return {
				Params: {},
				LPool: {},
				LPoolAll: {},
				LProviderRecord: {},
				EstimateSwapOut: {},
				EstimateSwapIn: {},
				EstimateContribution: {},
				MakeValidPair: {},
				EstimatePoolRemove: {},
				ListRecordsFromPool: {},
				
				_Structure: {
						LPool: getStructure(LPool.fromPartial({})),
						LProviderRecord: getStructure(LProviderRecord.fromPartial({})),
						Params: getStructure(Params.fromPartial({})),
						
		},
		_Registry: registry,
		_Subscriptions: new Set(),
	}
}

// initial state
const state = getDefaultState()

export default {
	namespaced: true,
	state,
	mutations: {
		RESET_STATE(state) {
			Object.assign(state, getDefaultState())
		},
		QUERY(state, { query, key, value }) {
			state[query][JSON.stringify(key)] = value
		},
		SUBSCRIBE(state, subscription) {
			state._Subscriptions.add(JSON.stringify(subscription))
		},
		UNSUBSCRIBE(state, subscription) {
			state._Subscriptions.delete(JSON.stringify(subscription))
		}
	},
	getters: {
				getParams: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.Params[JSON.stringify(params)] ?? {}
		},
				getLPool: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.LPool[JSON.stringify(params)] ?? {}
		},
				getLPoolAll: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.LPoolAll[JSON.stringify(params)] ?? {}
		},
				getLProviderRecord: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.LProviderRecord[JSON.stringify(params)] ?? {}
		},
				getEstimateSwapOut: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.EstimateSwapOut[JSON.stringify(params)] ?? {}
		},
				getEstimateSwapIn: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.EstimateSwapIn[JSON.stringify(params)] ?? {}
		},
				getEstimateContribution: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.EstimateContribution[JSON.stringify(params)] ?? {}
		},
				getMakeValidPair: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.MakeValidPair[JSON.stringify(params)] ?? {}
		},
				getEstimatePoolRemove: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.EstimatePoolRemove[JSON.stringify(params)] ?? {}
		},
				getListRecordsFromPool: (state) => (params = { params: {}}) => {
					if (!(<any> params).query) {
						(<any> params).query=null
					}
			return state.ListRecordsFromPool[JSON.stringify(params)] ?? {}
		},
				
		getTypeStructure: (state) => (type) => {
			return state._Structure[type].fields
		},
		getRegistry: (state) => {
			return state._Registry
		}
	},
	actions: {
		init({ dispatch, rootGetters }) {
			console.log('Vuex module: jackaldao.canine.lp initialized!')
			if (rootGetters['common/env/client']) {
				rootGetters['common/env/client'].on('newblock', () => {
					dispatch('StoreUpdate')
				})
			}
		},
		resetState({ commit }) {
			commit('RESET_STATE')
		},
		unsubscribe({ commit }, subscription) {
			commit('UNSUBSCRIBE', subscription)
		},
		async StoreUpdate({ state, dispatch }) {
			state._Subscriptions.forEach(async (subscription) => {
				try {
					const sub=JSON.parse(subscription)
					await dispatch(sub.action, sub.payload)
				}catch(e) {
					throw new Error('Subscriptions: ' + e.message)
				}
			})
		},
		
		
		
		 		
		
		
		async QueryParams({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JackaldaoCanineLp.query.queryParams()).data
				
					
				commit('QUERY', { query: 'Params', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryParams', payload: { options: { all }, params: {...key},query }})
				return getters['getParams']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryParams API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryLPool({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JackaldaoCanineLp.query.queryLPool( key.index)).data
				
					
				commit('QUERY', { query: 'LPool', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryLPool', payload: { options: { all }, params: {...key},query }})
				return getters['getLPool']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryLPool API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryLPoolAll({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JackaldaoCanineLp.query.queryLPoolAll(query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.JackaldaoCanineLp.query.queryLPoolAll({...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'LPoolAll', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryLPoolAll', payload: { options: { all }, params: {...key},query }})
				return getters['getLPoolAll']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryLPoolAll API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryLProviderRecord({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JackaldaoCanineLp.query.queryLProviderRecord( key.poolName,  key.provider)).data
				
					
				commit('QUERY', { query: 'LProviderRecord', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryLProviderRecord', payload: { options: { all }, params: {...key},query }})
				return getters['getLProviderRecord']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryLProviderRecord API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryEstimateSwapOut({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JackaldaoCanineLp.query.queryEstimateSwapOut( key.poolName,  key.inputCoin)).data
				
					
				commit('QUERY', { query: 'EstimateSwapOut', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryEstimateSwapOut', payload: { options: { all }, params: {...key},query }})
				return getters['getEstimateSwapOut']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryEstimateSwapOut API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryEstimateSwapIn({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JackaldaoCanineLp.query.queryEstimateSwapIn( key.poolName,  key.outputCoins)).data
				
					
				commit('QUERY', { query: 'EstimateSwapIn', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryEstimateSwapIn', payload: { options: { all }, params: {...key},query }})
				return getters['getEstimateSwapIn']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryEstimateSwapIn API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryEstimateContribution({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JackaldaoCanineLp.query.queryEstimateContribution( key.poolName,  key.desiredAmount)).data
				
					
				commit('QUERY', { query: 'EstimateContribution', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryEstimateContribution', payload: { options: { all }, params: {...key},query }})
				return getters['getEstimateContribution']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryEstimateContribution API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryMakeValidPair({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JackaldaoCanineLp.query.queryMakeValidPair( key.poolName,  key.coin)).data
				
					
				commit('QUERY', { query: 'MakeValidPair', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryMakeValidPair', payload: { options: { all }, params: {...key},query }})
				return getters['getMakeValidPair']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryMakeValidPair API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryEstimatePoolRemove({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JackaldaoCanineLp.query.queryEstimatePoolRemove( key.amount, query ?? undefined)).data
				
					
				while (all && (<any> value).pagination && (<any> value).pagination.next_key!=null) {
					let next_values=(await client.JackaldaoCanineLp.query.queryEstimatePoolRemove( key.amount, {...query ?? {}, 'pagination.key':(<any> value).pagination.next_key} as any)).data
					value = mergeResults(value, next_values);
				}
				commit('QUERY', { query: 'EstimatePoolRemove', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryEstimatePoolRemove', payload: { options: { all }, params: {...key},query }})
				return getters['getEstimatePoolRemove']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryEstimatePoolRemove API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		
		
		 		
		
		
		async QueryListRecordsFromPool({ commit, rootGetters, getters }, { options: { subscribe, all} = { subscribe:false, all:false}, params, query=null }) {
			try {
				const key = params ?? {};
				const client = initClient(rootGetters);
				let value= (await client.JackaldaoCanineLp.query.queryListRecordsFromPool( key.poolName)).data
				
					
				commit('QUERY', { query: 'ListRecordsFromPool', key: { params: {...key}, query}, value })
				if (subscribe) commit('SUBSCRIBE', { action: 'QueryListRecordsFromPool', payload: { options: { all }, params: {...key},query }})
				return getters['getListRecordsFromPool']( { params: {...key}, query}) ?? {}
			} catch (e) {
				throw new Error('QueryClient:QueryListRecordsFromPool API Node Unavailable. Could not perform query: ' + e.message)
				
			}
		},
		
		
		async sendMsgSwap({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JackaldaoCanineLp.tx.sendMsgSwap({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSwap:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgSwap:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgDepositLPool({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JackaldaoCanineLp.tx.sendMsgDepositLPool({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDepositLPool:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgDepositLPool:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgCreateLPool({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JackaldaoCanineLp.tx.sendMsgCreateLPool({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateLPool:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgCreateLPool:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		async sendMsgWithdrawLPool({ rootGetters }, { value, fee = [], memo = '' }) {
			try {
				const client=await initClient(rootGetters)
				const result = await client.JackaldaoCanineLp.tx.sendMsgWithdrawLPool({ value, fee: {amount: fee, gas: "200000"}, memo })
				return result
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgWithdrawLPool:Init Could not initialize signing client. Wallet is required.')
				}else{
					throw new Error('TxClient:MsgWithdrawLPool:Send Could not broadcast Tx: '+ e.message)
				}
			}
		},
		
		async MsgSwap({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JackaldaoCanineLp.tx.msgSwap({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgSwap:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgSwap:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgDepositLPool({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JackaldaoCanineLp.tx.msgDepositLPool({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgDepositLPool:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgDepositLPool:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgCreateLPool({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JackaldaoCanineLp.tx.msgCreateLPool({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgCreateLPool:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgCreateLPool:Create Could not create message: ' + e.message)
				}
			}
		},
		async MsgWithdrawLPool({ rootGetters }, { value }) {
			try {
				const client=initClient(rootGetters)
				const msg = await client.JackaldaoCanineLp.tx.msgWithdrawLPool({value})
				return msg
			} catch (e) {
				if (e == MissingWalletError) {
					throw new Error('TxClient:MsgWithdrawLPool:Init Could not initialize signing client. Wallet is required.')
				} else{
					throw new Error('TxClient:MsgWithdrawLPool:Create Could not create message: ' + e.message)
				}
			}
		},
		
	}
}
