package com.eagle.notification.bot.ui

import dev.inmo.tgbotapi.requests.abstracts.FileId
import dev.inmo.tgbotapi.types.ChatId

sealed class TelegramEvent {
    abstract val chatId: ChatId
}

data class EnterText(
    override val chatId: ChatId,
    val text: String,
    val replyId: Long?
) : TelegramEvent()

data class ButtonPress(
    override val chatId: ChatId,
    val index: Int,
    val replyId: Long
) : TelegramEvent()

data class SendPhoto(
    override val chatId: ChatId,
    val text: String?,
    val photos: List<FileId>
) : TelegramEvent()
