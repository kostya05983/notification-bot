package com.eagle.notification.bot.ui

import dev.inmo.tgbotapi.bot.TelegramBot
import dev.inmo.tgbotapi.extensions.api.send.sendMessage
import dev.inmo.tgbotapi.types.ChatId
import dev.inmo.tgbotapi.types.IdChatIdentifier
import dev.inmo.tgbotapi.types.MessageIdentifier
import dev.inmo.tgbotapi.types.message.HTMLParseMode

class Text(
    private val text: String
) : Element {
    override suspend fun publish(
        chatId: IdChatIdentifier,
        bot: TelegramBot
    ): MessageIdentifier {
        return bot.sendMessage(
            chatId = chatId,
            text = text,
            parseMode = HTMLParseMode
        ).messageId
    }
}
