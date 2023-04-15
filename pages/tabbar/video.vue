<template>
    <view class="video-view">
        <!-- 播放器 -->
        <video
            autoplay
            id="myVideo"
            class="video-content"
            src="../../static/video.mp4"
            :controls="false"
            @loadedmetadata="loadedmetadata"
            @timeupdate="timeUpdate"
        >
        </video>
        <!-- 遮盖层 -->
        <view class="video-cover" @click="touchPlay">
            <view class="view-cover-content">
                <!-- 信息栏 -->
                <view class="video-info">
                    <!-- 弹簧 -->
                    <view style="flex: 1;"></view>
                    <!-- 显示时间 -->
                    <view class="video-time" v-if="isDrag">
                        <view class="video-time-current">{{ dragStarTime }}</view>
                        <view style="padding: 0 20rpx;">/</view>
                        <text>{{ dragEndTime }}</text>
                    </view>
                    <!-- 进度栏 -->
                    <view class="slider-view" @click.stop>
                        <u-slider
                            v-model="currentTime"
                            min="0"
                            :max="duration"
                            inactiveColor="rgba(255, 255, 255, 0.2)"
                            activeColor="#F8DD52"
                            @changing="sliderChanging"
                            @change="sliderChange"
                        ></u-slider>
                    </view>
                    <!-- 当前集数 -->
                    <view class="video-info-current" @click.stop>
                        <!-- 名称 -->
                        <!-- <view class="video-info-title text-ell" v-if="drama.project_drama_name">{{ drama.project_drama_name }}-第{{ drama.eq_number }}集</view> -->
                        <!-- 弹簧 -->
                        <view style="flex: 1;"></view>
                        <!-- 选集 -->
                        <view class="video-info-btn" @click="touchSwitch">
                            <view class="arrow-right-title">选集</view>
                            <u-icon name="arrow-right" color="#fff"></u-icon>
                        </view>
                    </view>
                </view>
            </view>
        </view>
    </view>
</template>

<script>
export default {
    props: {
        // 当前剧集
        drama: {
            type: Object,
            default: () => {}
        }
    },
    data() {
        return {
            // 播放器上下文
            videoContext: undefined,
            // 播放状态
            isPlay: true,
            // 当前时长
            currentTime: 0,
            // 总时间
            duration: 0.1,
            // 是否正在拖拽进度
            isDrag: false,
            // 拖拽时视频时间
            dragStarTime: '',
            dragEndTime: '',
            // 当前还没显示过提示消息
            isShowMsg: true
        }
    },
    mounted () {
        // 获取播放器上下文（后面的 this 需要传入，在微信小程序上无法暂停播放拖拽精度，所以需要传入这个）
        this.videoContext = uni.createVideoContext('myVideo', this)
    },
    methods: {
        // 播放
        play () {
            this.videoContext.play()
        },
        // 暂停
        pause () {
            this.videoContext.pause()
        },
        // 播放状态切换
        touchPlay () {
            this.isPlay = !this.isPlay
            if (this.isPlay) {
                this.play ()
            } else {
                this.pause()
            }
        },
        // 播放进度发生变化
        timeUpdate (e) {
            // 拖拽时不需要进行更新
            if (!this.isDrag) {
                // 更新进度
                const { currentTime } = e.detail
                this.currentTime = currentTime
                // 播放完成
                if (Math.trunc(currentTime) === Math.trunc(duration)) {
                    this.$emit('playcomplete', e)
                }
                // 返回当前播放时间
                this.$emit('timeupdate', e)
            }
        },
        // 拖拽结束
        sliderChange (value) {
            // 停止拖拽
            this.isDrag = false
            // 判断一下是否大于基础时间
            if (this.duration > 0.1) {
                // 跳到指定时间点
                this.videoContext.seek(value)
                // 并调用播放
                this.play()
            }
        },
        // 正在拖拽
        sliderChanging (value) {
            // 开始拖拽
            this.isDrag = true
            // 刷新时间
            this.reloadVideoTime()
        },
        // 刷新显示时间
        reloadVideoTime () {
            // 当前时间
            this.dragStarTime = this.$pub.TIME_CONVERT(this.currentTime)
            // 总时间
            this.dragEndTime = this.$pub.TIME_CONVERT(this.duration)
        },
        // 加载完成
        loadedmetadata (e) {
            const { duration } = e.detail
            // 记录视频总时间
            this.duration = duration
            // 回调
            this.$emit('loadcomplete')
        },
        // 切换选集
        touchSwitch () {
            this.$emit('switch')
        }
    }
}
</script>

<style lang="scss">
.video-view {
    position: relative;
    width: 100%;
    height: 100%;
    .video-content {
        width: 100%;
        height: 100%;
    }
    .video-cover {
        position: absolute;
        left: 0;
        top: 0;
        width: 100%;
        height: 100%;
        .view-cover-content {
            position: relative;
            width: 100%;
            height: 100%;
            .video-info {
                display: flex;
                flex-direction: column;
                position: absolute;
                bottom: 0;
                left: 0;
                width: 100%;
                height: 246rpx;
                background: linear-gradient(180deg, rgba(0,0,0,0.00) 0%, rgba(0,0,0,0.00) 0%, rgba(0,0,0,0.40) 100%, rgba(0,0,0,0.40) 100%);
                .slider-view {
                    flex-shrink: 0;
                }
                .video-info-current {
                    flex-shrink: 0;
                    display: flex;
                    align-items: center;
                    width: 100%;
                    color: #fff;
                    font-size: 34rpx;
                    padding: 0 40rpx 40rpx;
                    width: calc(100% - 80rpx);
                    .video-info-btn {
                        flex-shrink: 0;
                        display: flex;
                        align-items: center;
                        padding-left: 40rpx;
                        .arrow-right-title {
                            margin-right: 10rpx;
                        }
                    }
                }
                .video-time {
                    display: flex;
                    align-items: center;
                    justify-content: center;
                    font-size: 64rpx;
                    color: rgba(255, 255, 255, 0.7);
                    .video-time-current {
                        color: #fff;
                    }
                }
            }
        }
    }
}
</style>
