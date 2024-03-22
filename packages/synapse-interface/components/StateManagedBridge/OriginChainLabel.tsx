import { useSelector } from 'react-redux'
import { ChevronDownIcon } from '@heroicons/react/outline'
import { CHAINS_BY_ID, ORDERED_CHAINS_BY_ID } from '@constants/chains'
import { getNetworkButtonBorder } from '@/styles/chains'
import Image from 'next/image'
import Tooltip from '@tw/Tooltip'
import { useEffect, useState } from 'react'
import { useDispatch } from 'react-redux'
import { setFromChainId } from '@/slices/bridgeSlice'
import { setShowFromChainSlideOver } from '@/slices/bridgeDisplaySlice'
import { RootState } from '@/store/store'
import { segmentAnalyticsEvent } from '@/contexts/SegmentAnalyticsProvider'

export const OriginChainLabel = ({
  chains,
  chainId,
  connectedChainId,
}: {
  chains: number[] | undefined
  chainId: number
  connectedChainId: number
}) => {
  const [orderedChains, setOrderedChains] = useState<number[]>([])
  useEffect(() => {
    setOrderedChains(chainOrderBySwapSide(chainId))
  }, [chainId, connectedChainId, chains])

  const dispatch = useDispatch()

  return (
    <div className="flex items-center justify-center md:justify-between">
      <div className={`text-gray-400 hidden md:block lg:block text-sm mr-2`}>
        Origin
      </div>
      <div className="flex items-center space-x-4 md:space-x-3">
        {orderedChains.map((id) =>
          Number(id) === chainId ? (
            <SelectedChain chainId={Number(id)} key={id} />
          ) : (
            <PossibleChain chainId={Number(id)} key={id} />
          )
        )}
        <button
          onClick={() => {
            dispatch(setShowFromChainSlideOver(true))
          }}
          tabIndex={0}
          data-test-id="bridge-origin-chain-list-button"
          className="w-8 h-8 px-1.5 py-1.5 bg-[#C4C4C4] bg-opacity-10 rounded-full hover:cursor-pointer group"
        >
          <ChevronDownIcon className="text-gray-300 transition transform-gpu group-hover:opacity-50 group-active:rotate-180" />
        </button>
      </div>
    </div>
  )
}

const PossibleChain = ({ chainId }: { chainId: number }) => {
  const chain = CHAINS_BY_ID[chainId]
  const { fromChainId } = useSelector((state: RootState) => state.bridge)

  const dispatch = useDispatch()

  const onChangeChain = () => {
    const eventTitle = `[Bridge User Action] Change Origin Chain`
    const eventData = {
      previousFromChainId: fromChainId,
      newFromChainId: chainId,
    }
    segmentAnalyticsEvent(eventTitle, eventData)
    dispatch(setFromChainId(chainId))
  }
  return chain ? (
    <button
      className="
        w-7 h-7
        md:w-7
        px-0.5 py-0.5
        border border-gray-500 rounded-full
      "
      tabIndex={0}
      onClick={onChangeChain}
    >
      <Tooltip content={chain.name}>
        <Image
          src={chain.chainImg}
          className="duration-300 rounded-full hover:scale-125"
          alt={chain.name}
        />
      </Tooltip>
    </button>
  ) : null
}

const SelectedChain = ({ chainId }: { chainId: number }) => {
  const chain = CHAINS_BY_ID[chainId]
  return chain ? (
    <div
      className={`
        px-1
        flex items-center
        bg-bgLight
        text-white
        border ${getNetworkButtonBorder(chain.color)}
        rounded-full
      `}
    >
      <Image
        alt="chain image"
        src={chain.chainImg}
        className="w-5 h-5 my-1 mr-0 rounded-full md:mr-1 opacity-80"
      />
      <div className="hidden md:inline-block lg:inline-block">
        <div className="mr-2 text-sm text-white">
          {chain.name === 'Boba Network' ? 'Boba' : chain.name}
        </div>
      </div>
    </div>
  ) : null
}

const chainOrderBySwapSide = (chainId: number) => {
  let orderedChains
  orderedChains = ORDERED_CHAINS_BY_ID.filter((e) => e !== String(chainId))
  orderedChains = orderedChains.slice(0, 5)
  orderedChains.unshift(chainId)

  return orderedChains
}