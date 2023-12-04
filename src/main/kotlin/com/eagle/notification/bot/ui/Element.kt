package com.eagle.notification.bot.ui

import dev.inmo.tgbotapi.bot.TelegramBot
import dev.inmo.tgbotapi.types.ChatId
import dev.inmo.tgbotapi.types.IdChatIdentifier
import dev.inmo.tgbotapi.types.MessageIdentifier

interface Element {
    suspend fun publish(
        chatId: IdChatIdentifier,
        bot: TelegramBot
    ): MessageIdentifier

    suspend fun onMessage(
        messageId: MessageIdentifier,
        event: TelegramEvent?,
        bot: TelegramBot
    ) = null
}
